package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type query struct {
	ArticleList []Article `xml:"artikel"`
}

type Article struct {
	Nr                    int            `xml:"nr"`
	ArticleID             int            `xml:"Artikelid"`
	ArticleNumber         int            `xml:"Varnummer"`
	Name                  string         `xml:"Namn"`
	SecondaryName         string         `xml:"Namn2"`
	PriceIncludingVAT     float32        `xml:"Prisinklmoms"`
	VolumeInMl            float32        `xml:"Volymiml"`
	PricePerLitre         float32        `xml:"PrisPerLiter"`
	SalesStart            salesStartDate `xml:"Saljstart"`
	Expired               bool           `xml:"Utgått"`
	ArticleGroup          string         `xml:"Varugrupp"`
	ArticleType           string         `xml:"Typ"`
	ArticleStyle          string         `xml:"Stil"`
	Packaging             string         `xml:"Forpackning"`
	Seal                  string         `xml:"Forslutning"`
	Origin                string         `xml:"Ursprung"`
	OriginCountry         string         `xml:"Ursprunglandnamn"`
	Producer              string         `xml:"Producent"`
	Supplier              string         `xml:"Leverantor"`
	Vintage               string         `xml:"Argang"`
	AlcoholPercentage     percent        `xml:"Alkoholhalt"`
	Selection             string         `xml:"Sortiment"`
	SelectionText         string         `xml:"SortimentText"`
	Organic               bool           `xml:"Ekologisk"`
	Ethical               bool           `xml:"Etiskt"`
	Koscher               bool           `xml:"Koscher"`
	IngredientDescription string         `xml:"RavarorBeskrivning"`
}

func (e Article) String() string {
	return fmt.Sprintf(`Nr: %02d 
                        ArticleID: %02d 
                        ArticleNumber: %02d 
                        Name: %s 
                        SecondaryName: %s 
                        PriceIncludingVAT: %f 
                        VolumeInMl: %f 
                        PricePerLitre: %f 
                        SalesStart: %s 
                        Expired: %t
                        ArticleGroup: %s
                        ArticleType: %s
                        ArticleStyle: %s
                        Packaging: %s,
                        Seal: %s,
                        Origin: %s,
                        OriginCountry: %s,
                        Producer: %s,
                        Supplier: %s,
                        Vintage: %s,
                        AlcoholPercentage: %f,
                        Selection: %s,
						SelectionText: %s,
                        Organic: %t,
                        Ethical: %t,
                        Koscher: %t,
                        IngredientDescription: %s`,
		e.Nr,
		e.ArticleID,
		e.ArticleNumber,
		e.Name,
		e.SecondaryName,
		e.PriceIncludingVAT,
		e.VolumeInMl,
		e.PricePerLitre,
		e.SalesStart,
		e.Expired,
		e.ArticleGroup,
		e.ArticleType,
		e.ArticleStyle,
		e.Packaging,
		e.Seal,
		e.Origin,
		e.OriginCountry,
		e.Producer,
		e.Supplier,
		e.Vintage,
		e.AlcoholPercentage,
		e.Selection,
		e.SelectionText,
		e.Organic,
		e.Ethical,
		e.Koscher,
		e.IngredientDescription)
}

type salesStartDate struct {
	time.Time
}

func (date *salesStartDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02" // yyyymmdd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*date = salesStartDate{parse}
	return nil
}

type percent struct {
	float64
}

func (p *percent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := strconv.ParseFloat(strings.Trim(v, "%"), 64)

	if err != nil {
		fmt.Println("Error parsing percent:", err)
	}

	*p = percent{parse}
	return nil
}

func (a Article) ParseArticleData(data []byte) ([]Article, error) {
	var q query
	err := xml.Unmarshal(data, &q)

	if err != nil {
		return nil, err
	}

	return q.ArticleList, nil
}
