package pkg

import (
	"Reach-page/internal/pkg/RPC/pb"
	"Reach-page/internal/pkg/database"
	"context"
	"log"
)

type PageServer struct {
	pb.PagePackageServer
}

func (s *PageServer) GetPage(ctx context.Context, input *pb.IdRequest) (*pb.Page, error) {
	res := database.GetPage(input.Id)
	template := DataBaseTemplateToRPCTemplate(res.Template)
	links := []*pb.PageLinks{}
	for _, el := range res.Links {
		link := DatabaseLinkToRPCLink(el)
		links = append(links, &link)
	}
	return &pb.Page{
		Id:            res.Id,
		Route:         res.Route,
		Template:      &template,
		UserAccountId: res.UserAccountId,
		Links:         links,
	}, nil
}

func (s *PageServer) CreatePage(ctx context.Context, input *pb.PageRequest) (*pb.PageResponse, error) {
	res := database.CreatePage(input.UserAccountId, input.Route)
	return &pb.PageResponse{
		Id: res,
	}, nil
}

func (s *PageServer) CreateTemplate(ctx context.Context, input *pb.TemplateRequest) (*pb.Template, error) {
	res := database.CreateTemplate(
		input.PageId,
		input.Name,
		input.Desc,
		input.Image,
		input.Button,
		input.Background,
		input.Font,
		input.FontColor,
		input.Social,
		input.SocialPosition,
	)
	rpcTemplate := DataBaseTemplateToRPCTemplate(res)
	return &rpcTemplate, nil
}

func (s *PageServer) UpdateTemplate(ctx context.Context, input *pb.TemplateRequest) (*pb.VoidResponse, error) {
	values := database.Template{
		Name:           input.Name,
		Desc:           input.Desc,
		Image:          input.Image,
		Button:         input.Button,
		Background:     input.Background,
		Font:           input.Font,
		FontColor:      input.FontColor,
		SocialPosition: input.SocialPosition,
	}
	database.UpdateTemplate(
		input.PageId, values,
	)
	return &pb.VoidResponse{}, nil
}

func (s *PageServer) CreateLink(ctx context.Context, input *pb.CreateLinkRequest) (*pb.PageLinks, error) {
	res := database.CreateLink(input.PageId, input.Name, input.Link, input.Icon, input.IsSocialIcon, int(input.Sequence))
	rpcLink := DatabaseLinkToRPCLink(res)
	return &rpcLink, nil
}

func (s *PageServer) UpdateLink(ctx context.Context, links *pb.PageLinks) (*pb.VoidResponse, error) {
	values := database.PageLinks{
		Id:     links.Id,
		Name:   links.Name,
		Link:   links.Link,
		Icon:   links.Icon,
		Social: links.Social,
		PageId: links.PageId,
	}
	database.UpdateLink(links.Id, values)
	return &pb.VoidResponse{}, nil
}

func (s *PageServer) CreateMetaLink(ctx context.Context, input *pb.Meta) (*pb.Meta, error) {
	res := database.CreateMetaLinks(input.TemplateId, input.Type, input.Value)
	rpcMeta := DatabaseMetaToRPCMeta(res)
	return &rpcMeta, nil
}

func (s *PageServer) UpdateMetaLink(ctx context.Context, input *pb.Meta) (*pb.VoidResponse, error) {
	values := database.Meta{
		Id:         input.Id,
		Value:      input.Value,
		Type:       input.Type,
		TemplateId: input.TemplateId,
	}
	database.UpdateMetaLink(input.Id, values)
	return &pb.VoidResponse{}, nil
}

func (s *PageServer) GetPageId(ctx context.Context, input *pb.IdRequest) (*pb.Page, error) {
	log.Println("Input page id")
	log.Println(input.Id)
	res := database.GetPageId(input.Id)
	template := DataBaseTemplateToRPCTemplate(res.Template)
	links := []*pb.PageLinks{}
	for _, el := range res.Links {
		link := DatabaseLinkToRPCLink(el)
		links = append(links, &link)
	}
	return &pb.Page{
		Id:            res.Id,
		Route:         res.Route,
		Template:      &template,
		UserAccountId: res.UserAccountId,
		Links:         links,
	}, nil
}
