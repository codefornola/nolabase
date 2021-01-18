package assessor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cridenour/go-postgis"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func ParseProperty(doc *goquery.Document) *Property {
	propTable := tableWithTitle(doc.Find("table"), "OWNER AND PARCEL INFORMATION")
	headers := propTable.Find("tbody > tr > td.owner_header").Map(func(i int, selection *goquery.Selection) string {
		return strings.ToUpper(strings.TrimSpace(selection.Text()))
	})
	property := &Property{}
	propTable.Find("tbody > tr > td.owner_value").Each(func(i int, selection *goquery.Selection) {
		value := strings.TrimSpace(selection.Text())
		header := headers[i]
		switch header {
		case "OWNER NAME":
			property.OwnerName = value
		case "TAX BILL NUMBER":
			property.TaxBillNumber = value
		case "LOCATION ADDRESS":
			property.LocationAddress = value
		case "MAILING ADDRESS":
			property.MailingAddress = value
		case "PROPERTY CLASS":
			property.PropertyClass = strings.ToUpper(value)
		case "ASSESSMENT AREA":
			property.AssessmentArea = strings.ToUpper(value)
		case "LAND AREA (SQ FT)":
			property.LandAreaSqFt = parseInt(value)
		case "BUILDING AREA (SQ FT)":
			property.BuildingAreaSqFt = parseInt(value)
		case "PARCEL MAP":
			if href, ok := selection.Find("a").Attr("href"); ok {
				u, err := url.Parse(href)
				if err != nil {
					log.Fatal(err)
				}
				extent := u.Query().Get("extent")
				if extent == "" {
					log.Fatal("Can't find extent")
				}
				parts := strings.Split(extent, " ")
				x1, _ := strconv.ParseInt(parts[0], 10, 64)
				y1, _ := strconv.ParseInt(parts[1], 10, 64)
				x2, _ := strconv.ParseInt(parts[2], 10, 64)
				y2, _ := strconv.ParseInt(parts[3], 10, 64)
				x := float64(x1) + (float64(x2 - x1) / 2)
				y := float64(y1) + (float64(y2 - y1) / 2)
				property.LngLatPoint = postgis.PointS{SRID: 3452, X: x, Y: y}
			} else {
				log.Printf("Couldn't find map href %s %+v\n", header, selection)
			}
		default:
			//log.Printf("Don't know how to map property value %s => %s\n", header, value)
		}
	})

	return property
}

func ParseValues(doc *goquery.Document) []*PropertyValue {
	valuesTable := tableWithTitle(doc.Find("table"), "VALUE INFORMATION")
	headers := valuesTable.Find("tbody > tr > td.tax_header").Map(func(i int, s *goquery.Selection) string {
		return strings.ToUpper(strings.TrimSpace(s.Text()))
	})
	var values []*PropertyValue
	valuesTable.Find("tr").Each(func(i int, s *goquery.Selection) {
		taxVals := s.Find("td.tax_value")
		// does it have any of these?
		if taxVals.Length() > 0 {
			v := &PropertyValue{}
			taxVals.Each(func(j int, n *goquery.Selection) {
				value := strings.TrimSpace(n.Text())
				header := headers[j]
				switch header {
				case "YEAR":
					v.Year = parseInt(value)
				case "LAND VALUE":
					v.LandValue = parseInt(value)
				case "BUILDING VALUE":
					v.BuildingValue = parseInt(value)
				case "TOTAL VALUE":
					v.TotalValue = parseInt(value)
				case "ASSESSEDLAND VALUE":
					v.AssessedLandValue = parseInt(value)
				case "ASSESSEDBUILDING VALUE":
					v.AssessedBuildingValue = parseInt(value)
				case "TOTALASSESSED VALUE":
					v.TotalAssessedValue = parseInt(value)
				case "HOMESTEADEXEMPTION VALUE":
					v.HomesteadExemptionValue = parseInt(value)
				case "TAXABLEASSESSMENT":
					v.TaxableAssessment = parseInt(value)
				case "AGE FREEZE":
					v.AgeFreeze = parseInt(value)
				case "DISABILITY FREEZE":
					v.DisabilityFreeze = parseInt(value)
				case "ASSMNT CHANGE":
					v.AssessmentChange = parseInt(value)
				case "TAX CONTRACT":
					v.TaxContract = parseInt(value)
				default:
					//log.Printf("Don't know how to map tax value %s => %s\n", header, value)
				}
			})
			values = append(values, v)
		}
	})

	return values
}

func ParseSales(doc *goquery.Document) []*PropertySale {
	salesTable := tableWithTitle(doc.Find("table"), "SALE/TRANSFER INFORMATION")
	headers := salesTable.Find("tbody > tr > td.sales_header").Map(func(i int, s *goquery.Selection) string {
		return strings.ToUpper(strings.TrimSpace(s.Text()))
	})
	var sales []*PropertySale
	salesTable.Find("tr").Each(func(i int, s *goquery.Selection) {
		salesVals := s.Find("td.sales_value")
		// does it have any of these?
		if salesVals.Length() > 0 {
			v := &PropertySale{}
			salesVals.Each(func(j int, n *goquery.Selection) {
				value := strings.TrimSpace(n.Text())
				header := headers[j]
				switch header {
				case "SALE/TRANSFER DATE":
					v.Date = value
				case "GRANTOR":
					v.Grantor = value
				case "GRANTEE":
					v.Grantee = value
				case "NOTARIAL ARCHIVE NUMBER":
					v.NotarialArchiveNumber = value
				case "INSTRUMENT NUMBER":
					v.InstrumentNumber = value
				case "PRICE":
					v.Price = parseInt(value)
				default:
					//log.Printf("Don't know how to map sales value %s => %s\n", header, value)
				}
			})
			sales = append(sales, v)
		}
	})

	return sales
}

func tableWithTitle(node *goquery.Selection, title string) *goquery.Selection {
	var selection goquery.Selection
	node.Each(func(i int, s *goquery.Selection) {
		el := s.Find("tr > td.table_header").First()
		if strings.ToUpper(strings.TrimSpace(el.Text())) == title {
			selection = *s
		}
	})
	return &selection
}

func parseInt(s string) *int {
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, " ", "")
	z, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		//log.Printf("Failed to parse invalid integer %s", s)
	}
	i := int(z)
	return &i
}

