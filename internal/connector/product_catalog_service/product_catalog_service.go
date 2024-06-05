package product_catalog_service

import (
	"context"
	product_service "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
)

type productCatalogServiceClient struct {
	client product_service.ProductCatalogServiceClient
}

func NewProductCatalogServiceClient(client product_service.ProductCatalogServiceClient) *productCatalogServiceClient {
	return &productCatalogServiceClient{
		client: client,
	}
}

func (c *productCatalogServiceClient) GetById(ctx context.Context, id []string) (*product_service.GetProductByIdResponse, error) {
	result, err := c.client.GetProductById(ctx, &product_service.GetProductByIdRequest{
		Ids: id,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
