package pkg

import (
	"Reach-page/internal/pkg/RPC/pb"
	"Reach-page/internal/pkg/database"
)

func DataBaseTemplateToRPCTemplate(template database.Template) pb.Template {
	metas := []*pb.Meta{}
	for _, tag := range template.MetaTags {
		meta := DatabaseMetaToRPCMeta(tag)
		metas = append(metas, &meta)
	}
	return pb.Template{
		Id:             template.Id,
		Name:           template.Name,
		Desc:           template.Desc,
		Image:          template.Image,
		Button:         template.Button,
		Background:     template.Background,
		Font:           template.Font,
		FontColor:      template.FontColor,
		MetaTags:       metas,
		PageId:         template.PageId,
		SocialPosition: template.SocialPosition,
	}
}

func DatabaseMetaToRPCMeta(tag database.Meta) pb.Meta {
	return pb.Meta{
		Id:         tag.Id,
		Value:      tag.Value,
		Type:       tag.Type,
		TemplateId: tag.TemplateId,
	}
}

func DatabaseLinkToRPCLink(link database.PageLinks) pb.PageLinks {
	return pb.PageLinks{
		Id:     link.Id,
		Name:   link.Name,
		Link:   link.Link,
		Icon:   link.Icon,
		Social: link.Social,
		PageId: link.PageId,
	}
}
