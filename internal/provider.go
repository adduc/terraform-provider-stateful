package internal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &StatefulProvider{}

type StatefulProvider struct {
	version string
}

type StatefulProviderModel struct {
	State types.Dynamic `tfsdk:"state"`
}

func (p *StatefulProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "stateful"
	resp.Version = p.version
}

func (p *StatefulProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"state": schema.DynamicAttribute{
				MarkdownDescription: "State to be stored in the provider",
				Required:            true,
			},
		},
	}
}

func (p *StatefulProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data StatefulProviderModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.DataSourceData = data.State
}

func (p *StatefulProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExpressionDataSource,
	}
}

func (p *StatefulProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &StatefulProvider{
			version: version,
		}
	}
}
