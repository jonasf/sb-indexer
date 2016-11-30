package main

import "testing"

func TestParseData(t *testing.T) {
	article := &Article{}
	rawdata := `<?xml version="1.0" encoding="utf-8"?>
        <artiklar xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
            <artikel>
            <nr>3052603</nr>
            <Artikelid>963301</Artikelid>
            <Varnummer>30526</Varnummer>
            <Namn>Räntmästarens Röda</Namn>
            <Namn2>Irländsk röd lager</Namn2>
            <Prisinklmoms>35.90</Prisinklmoms>
            <Volymiml>330.00</Volymiml>
            <PrisPerLiter>108.79</PrisPerLiter>
            <Saljstart>2015-06-01</Saljstart>
            <Utgått>0</Utgått>
            <Varugrupp>Öl</Varugrupp>
            <Typ>Mellanmörk lager</Typ>
            <Stil>Märzen och wienerstil</Stil>
            <Forpackning>Flaska</Forpackning>
            <Forslutning />
            <Ursprung>Skåne län</Ursprung>
            <Ursprunglandnamn>Sverige</Ursprunglandnamn>
            <Producent>Hönsinge Hantwerksbryggeri</Producent>
            <Leverantor>Hönsinge Hantwerksbryggeri AB</Leverantor>
            <Argang>2015</Argang>
            <Provadargang />
            <Alkoholhalt>5.50%</Alkoholhalt>
            <Sortiment>TSLS</Sortiment>
            <SortimentText>Lokalt och småskaligt</SortimentText>
            <Ekologisk>0</Ekologisk>
            <Etiskt>0</Etiskt>
            <Koscher>0</Koscher>
            <RavarorBeskrivning>Pilsner-, munich- och karamellmalt samt humle av sorterna perle, citra och cascade.</RavarorBeskrivning>
        </artikel>
        </artiklar>`

	result, err := article.ParseArticleData([]byte(rawdata))

	if err != nil {
		t.Errorf("Expected error not be nil but got %q", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected result count to be 1 but got %q", len(result))
	}

	a := result[0]

	if a.Nr != 3052603 {
		t.Errorf("Expected Nr to be 3052603 but got %q", a.Nr)
	}
	if a.ArticleID != 963301 {
		t.Errorf("Expected ArticleID to be 963301 but got %q", a.ArticleID)
	}
	if a.ArticleNumber != 30526 {
		t.Errorf("Expected ArticleNumber to be 30526 but got %q", a.ArticleNumber)
	}
	if a.Name != "Räntmästarens Röda" {
		t.Errorf("Expected Name to be Räntmästarens Röda but got %q", a.Name)
	}
	if a.SecondaryName != "Irländsk röd lager" {
		t.Errorf("Expected SecondaryName to be Irländsk röd lager but got %q", a.SecondaryName)
	}
	if a.PriceIncludingVAT != 35.90 {
		t.Errorf("Expected PriceIncludingVAT to be 35.90 but got %q", a.PriceIncludingVAT)
	}
	if a.VolumeInMl != 330.00 {
		t.Errorf("Expected VolumeInMl to be 330.00 but got %q", a.VolumeInMl)
	}
	if a.PricePerLitre != 108.79 {
		t.Errorf("Expected PricePerLitre to be 108.79 but got %q", a.PricePerLitre)
	}
	if a.SalesStart.Year() != 2015 && a.SalesStart.Month() != 6 && a.SalesStart.Day() != 1 {
		t.Errorf("Expected SalesStart to be 2015-06-01 but got %q", a.SalesStart)
	}
	if a.Expired != false {
		t.Errorf("Expected Expired to be false but got %q", a.Expired)
	}
	if a.ArticleGroup != "Öl" {
		t.Errorf("Expected ArticleGroup to be Öl but got %q", a.ArticleGroup)
	}
	if a.ArticleType != "Mellanmörk lager" {
		t.Errorf("Expected ArticleType to be Mellanmörk lager but got %q", a.ArticleType)
	}
	if a.ArticleStyle != "Märzen och wienerstil" {
		t.Errorf("Expected ArticleStyle to be Märzen och wienerstil but got %q", a.ArticleStyle)
	}
	if a.Packaging != "Flaska" {
		t.Errorf("Expected Packaging to be Flaska but got %q", a.Packaging)
	}
	if a.Seal != "" {
		t.Errorf("Expected Seal to be empty but got %q", a.Seal)
	}
	if a.Origin != "Skåne län" {
		t.Errorf("Expected Origin to be Skåne län but got %q", a.Origin)
	}
	if a.OriginCountry != "Sverige" {
		t.Errorf("Expected OriginCountry to be Sverige but got %q", a.OriginCountry)
	}
	if a.Producer != "Hönsinge Hantwerksbryggeri" {
		t.Errorf("Expected Producer to be Hönsinge Hantwerksbryggeri but got %q", a.Producer)
	}
	if a.Supplier != "Hönsinge Hantwerksbryggeri AB" {
		t.Errorf("Expected Supplier to be Hönsinge Hantwerksbryggeri AB but got %q", a.Supplier)
	}
	if a.Vintage != "2015" {
		t.Errorf("Expected Vintage to be 2015 but got %q", a.Vintage)
	}
	if a.AlcoholPercentage.float64 != 5.50 {
		t.Errorf("Expected AlcoholPercentage to be 5.50 but got %q", a.AlcoholPercentage)
	}
	if a.Selection != "TSLS" {
		t.Errorf("Expected Selection to be TSLS but got %q", a.Selection)
	}
	if a.SelectionText != "Lokalt och småskaligt" {
		t.Errorf("Expected SelectionText to be Lokalt och småskaligt but got %q", a.SelectionText)
	}
	if a.Organic != false {
		t.Errorf("Expected Organic to be false but got %q", a.Organic)
	}
	if a.Ethical != false {
		t.Errorf("Expected Ethical to be false but got %q", a.Ethical)
	}
	if a.Koscher != false {
		t.Errorf("Expected Koscher to be false but got %q", a.Koscher)
	}
	if a.IngredientDescription != "Pilsner-, munich- och karamellmalt samt humle av sorterna perle, citra och cascade." {
		t.Errorf("Expected IngredientDescription to be Pilsner-, munich- och karamellmalt samt humle av sorterna perle, citra och cascade. but got %q", a.IngredientDescription)
	}
}

func TestEmptyVintageYear(t *testing.T) {
	article := &Article{}
	rawdata := `<?xml version="1.0" encoding="utf-8"?>
        <artiklar xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
            <artikel>
            <nr>3052603</nr>
            <Artikelid>963301</Artikelid>
            <Varnummer>30526</Varnummer>
            <Namn>Räntmästarens Röda</Namn>
            <Namn2>Irländsk röd lager</Namn2>
            <Prisinklmoms>35.90</Prisinklmoms>
            <Volymiml>330.00</Volymiml>
            <PrisPerLiter>108.79</PrisPerLiter>
            <Saljstart>2015-06-01</Saljstart>
            <Utgått>0</Utgått>
            <Varugrupp>Öl</Varugrupp>
            <Typ>Mellanmörk lager</Typ>
            <Stil>Märzen och wienerstil</Stil>
            <Forpackning>Flaska</Forpackning>
            <Forslutning />
            <Ursprung>Skåne län</Ursprung>
            <Ursprunglandnamn>Sverige</Ursprunglandnamn>
            <Producent>Hönsinge Hantwerksbryggeri</Producent>
            <Leverantor>Hönsinge Hantwerksbryggeri AB</Leverantor>
            <Argang />
            <Provadargang />
            <Alkoholhalt>5.50%</Alkoholhalt>
            <Sortiment>TSLS</Sortiment>
            <SortimentText>Lokalt och småskaligt</SortimentText>
            <Ekologisk>0</Ekologisk>
            <Etiskt>0</Etiskt>
            <Koscher>0</Koscher>
            <RavarorBeskrivning>Pilsner-, munich- och karamellmalt samt humle av sorterna perle, citra och cascade.</RavarorBeskrivning>
        </artikel>
        </artiklar>`

	result, err := article.ParseArticleData([]byte(rawdata))

	if err != nil {
		t.Errorf("Expected error not be nil but got %q", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected result count to be 1 but got %q", len(result))
	}

	a := result[0]

	if a.Vintage != "" {
		t.Errorf("Expected Vintage to be 0 but got %q", a.Vintage)
	}
}
