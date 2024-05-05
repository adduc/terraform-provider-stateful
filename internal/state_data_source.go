package internal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigure = &ExpressionDataSource{}

func NewExpressionDataSource() datasource.DataSource {
	return &ExpressionDataSource{}
}

type ExpressionDataSource struct {
	state types.Dynamic `tfsdk:"state"`
}

func (d *ExpressionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(types.Dynamic)

	if !ok {
		resp.Diagnostics.AddError("invalid provider data", "")
		return
	}

	d.state = data
}

type ExpressionDataSourceModel struct {
	State types.Dynamic `tfsdk:"state"`
}

func (d *ExpressionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_state"
}

func (d *ExpressionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"state": schema.DynamicAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ExpressionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var data ExpressionDataSourceModel
	data.State = d.state

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
