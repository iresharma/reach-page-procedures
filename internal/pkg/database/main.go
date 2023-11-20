package database

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Page struct {
	gorm.Model
	Id            string
	Route         string
	Template      Template
	UserAccountId string
	Links         []PageLinks
}

type Template struct {
	gorm.Model
	Id             string
	Name           string
	Desc           string
	Image          string
	Button         string
	Background     string
	Font           string
	FontColor      string
	MetaTags       []Meta
	PageId         string
	SocialPosition string
}

type Meta struct {
	gorm.Model
	Id         string
	Value      string
	Type       string
	TemplateId string
}

type PageLinks struct {
	gorm.Model
	Id       string
	Name     string
	Link     string
	Icon     string
	Social   bool
	PageId   string
	Sequence int
}

var connStr = os.Getenv("POSTGRES")
var DB *gorm.DB = nil

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Cannot connect to database")
	}

	err = db.AutoMigrate(&Page{}, &Template{}, &Meta{}, &PageLinks{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}

func ServerBuild() []string {
	var routes []string
	if err := DB.Table("pages").Distinct().Pluck("route", &routes).Error; err != nil {
		log.Println("fetch routes failed")
	}
	return routes
}

func CreatePage(userAccountId string, route string) string {
	pageId := uuid.New()
	page := &Page{Id: pageId.String(), UserAccountId: userAccountId, Route: route}
	if err := DB.Create(&page).Error; err != nil {
		log.Println("Fataaa")
	}
	if err := DB.Table("user_accounts").Where("id = ?", userAccountId).Update("page_id", pageId).Error; err != nil {
		log.Println(err)
		log.Println("Update fata")
	}
	return page.Id
}

func GetPage(route string) Page {
	var result Page
	var template Template
	var metas []Meta
	var links []PageLinks
	if err := DB.First(&result, "route = ?", route).Error; err != nil {
		log.Println(err)
	}
	if err := DB.First(&template, "page_id = ?", result.Id).Error; err != nil {
		log.Println(err)
	}
	if err := DB.Where("template_id = ?", template.Id).Find(&metas).Error; err != nil {
		log.Println(err)
	}
	if err := DB.Where("page_id = ?", result.Id).Find(&links).Order("sequence").Error; err != nil {
		log.Println(err)
	}
	result.Links = links
	template.MetaTags = metas
	result.Template = template
	return result
}

func GetPageId(route string) Page {
	var result Page
	var template Template
	var metas []Meta
	var links []PageLinks
	if err := DB.First(&result, "id = ?", route).Error; err != nil {
		log.Println(err)
	}
	if err := DB.First(&template, "page_id = ?", result.Id).Error; err != nil {
		log.Println(err)
	}
	if err := DB.Where("template_id = ?", template.Id).Find(&metas).Error; err != nil {
		log.Println(err)
	}
	if err := DB.Where("page_id = ?", result.Id).Find(&links).Order("sequence").Error; err != nil {
		log.Println(err)
	}
	result.Links = links
	template.MetaTags = metas
	result.Template = template
	return result
}

func CreateTemplate(pageId string, Name string, Desc string, Image string, Button string, Background string, Font string, FontColor string, Social bool, SocialPosition string) Template {
	templateId := uuid.New().String()
	template := Template{
		Id:             templateId,
		Name:           Name,
		Desc:           Desc,
		Image:          Image,
		Button:         Button,
		Background:     Background,
		Font:           Font,
		FontColor:      FontColor,
		MetaTags:       nil,
		PageId:         pageId,
		SocialPosition: SocialPosition,
	}
	if err := DB.Create(&template).Error; err != nil {
		log.Println("wassup bitch i knew this will happen")
	}
	return template
}

func UpdateTemplate(pageId string, values Template) {
	if err := DB.Model(&Template{}).Where("page_id = ?", pageId).Updates(values).Error; err != nil {
		log.Println(err)
	}
}

func CreateLink(pageId string, Name string, Link string, icon string, social bool, sequence int) PageLinks {
	linkId := uuid.New().String()
	pageLink := PageLinks{Id: linkId, PageId: pageId, Name: Name, Link: Link, Icon: icon, Social: social, Sequence: sequence}
	if err := DB.Create(&pageLink).Error; err != nil {
		log.Println(err)
	}
	return pageLink
}

func UpdateLink(linkId string, vales PageLinks) {
	if err := DB.Model(&PageLinks{}).Where("id = ?", linkId).Updates(vales).Error; err != nil {
		log.Println(err)
	}
}

func CreateMetaLinks(templateId string, tagType string, value string) Meta {
	metaId := uuid.New().String()
	meta := Meta{Id: metaId, Type: tagType, Value: value, TemplateId: templateId}
	if err := DB.Create(&meta).Error; err != nil {
		log.Println(err)
	}
	return meta
}

func UpdateMetaLink(metaId string, values Meta) {
	if err := DB.Model(&Meta{}).Where("id = ?", metaId).Updates(values).Error; err != nil {
		log.Println(err)
	}
}
