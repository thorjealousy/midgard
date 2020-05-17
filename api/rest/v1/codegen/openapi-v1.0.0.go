// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AssetDetail defines model for AssetDetail.
type AssetDetail struct {
	Asset       *Asset  `json:"asset,omitempty"`
	DateCreated *int64  `json:"dateCreated,omitempty"`
	PriceRune   *string `json:"priceRune,omitempty"`
}

// BlockRewards defines model for BlockRewards.
type BlockRewards struct {
	BlockReward *string `json:"blockReward,omitempty"`
	BondReward  *string `json:"bondReward,omitempty"`
	StakeReward *string `json:"stakeReward,omitempty"`
}

// BondMetrics defines model for BondMetrics.
type BondMetrics struct {
	AverageActiveBond  *string `json:"averageActiveBond,omitempty"`
	AverageStandbyBond *string `json:"averageStandbyBond,omitempty"`
	MaximumActiveBond  *string `json:"maximumActiveBond,omitempty"`
	MaximumStandbyBond *string `json:"maximumStandbyBond,omitempty"`
	MedianActiveBond   *string `json:"medianActiveBond,omitempty"`
	MedianStandbyBond  *string `json:"medianStandbyBond,omitempty"`
	MinimumActiveBond  *string `json:"minimumActiveBond,omitempty"`
	MinimumStandbyBond *string `json:"minimumStandbyBond,omitempty"`
	TotalActiveBond    *string `json:"totalActiveBond,omitempty"`
	TotalStandbyBond   *string `json:"totalStandbyBond,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// NetworkInfo defines model for NetworkInfo.
type NetworkInfo struct {
	ActiveBonds             *[]string     `json:"activeBonds,omitempty"`
	ActiveNodeCount         *int          `json:"activeNodeCount,omitempty"`
	BlockRewards            *BlockRewards `json:"blockRewards,omitempty"`
	BondMetrics             *BondMetrics  `json:"bondMetrics,omitempty"`
	BondingROI              *string       `json:"bondingROI,omitempty"`
	NextChurnHeight         *string       `json:"nextChurnHeight,omitempty"`
	PoolActivationCountdown *int64        `json:"poolActivationCountdown,omitempty"`
	PoolShareFactor         *string       `json:"poolShareFactor,omitempty"`
	StakingROI              *string       `json:"stakingROI,omitempty"`
	StandbyBonds            *[]string     `json:"standbyBonds,omitempty"`
	StandbyNodeCount        *int          `json:"standbyNodeCount,omitempty"`
	TotalReserve            *string       `json:"totalReserve,omitempty"`
	TotalStaked             *string       `json:"totalStaked,omitempty"`
}

// NodeKey defines model for NodeKey.
type NodeKey struct {
	Ed25519   *string `json:"ed25519,omitempty"`
	Secp256k1 *string `json:"secp256k1,omitempty"`
}

// PoolDetail defines model for PoolDetail.
type PoolDetail struct {
	Asset            *Asset  `json:"asset,omitempty"`
	AssetDepth       *string `json:"assetDepth,omitempty"`
	AssetROI         *string `json:"assetROI,omitempty"`
	AssetStakedTotal *string `json:"assetStakedTotal,omitempty"`
	BuyAssetCount    *string `json:"buyAssetCount,omitempty"`
	BuyFeeAverage    *string `json:"buyFeeAverage,omitempty"`
	BuyFeesTotal     *string `json:"buyFeesTotal,omitempty"`
	BuySlipAverage   *string `json:"buySlipAverage,omitempty"`
	BuyTxAverage     *string `json:"buyTxAverage,omitempty"`
	BuyVolume        *string `json:"buyVolume,omitempty"`
	PoolDepth        *string `json:"poolDepth,omitempty"`
	PoolFeeAverage   *string `json:"poolFeeAverage,omitempty"`
	PoolFeesTotal    *string `json:"poolFeesTotal,omitempty"`
	PoolROI          *string `json:"poolROI,omitempty"`
	PoolROI12        *string `json:"poolROI12,omitempty"`
	PoolSlipAverage  *string `json:"poolSlipAverage,omitempty"`
	PoolStakedTotal  *string `json:"poolStakedTotal,omitempty"`
	PoolTxAverage    *string `json:"poolTxAverage,omitempty"`
	PoolUnits        *string `json:"poolUnits,omitempty"`
	PoolVolume       *string `json:"poolVolume,omitempty"`
	PoolVolume24hr   *string `json:"poolVolume24hr,omitempty"`
	Price            *string `json:"price,omitempty"`
	RuneDepth        *string `json:"runeDepth,omitempty"`
	RuneROI          *string `json:"runeROI,omitempty"`
	RuneStakedTotal  *string `json:"runeStakedTotal,omitempty"`
	SellAssetCount   *string `json:"sellAssetCount,omitempty"`
	SellFeeAverage   *string `json:"sellFeeAverage,omitempty"`
	SellFeesTotal    *string `json:"sellFeesTotal,omitempty"`
	SellSlipAverage  *string `json:"sellSlipAverage,omitempty"`
	SellTxAverage    *string `json:"sellTxAverage,omitempty"`
	SellVolume       *string `json:"sellVolume,omitempty"`
	StakeTxCount     *string `json:"stakeTxCount,omitempty"`
	StakersCount     *string `json:"stakersCount,omitempty"`
	StakingTxCount   *string `json:"stakingTxCount,omitempty"`
	Status           *string `json:"status,omitempty"`
	SwappersCount    *string `json:"swappersCount,omitempty"`
	SwappingTxCount  *string `json:"swappingTxCount,omitempty"`
	WithdrawTxCount  *string `json:"withdrawTxCount,omitempty"`
}

// ScannerStatus defines model for ScannerStatus.
type ScannerStatus struct {
	Chain       *string `json:"chain,omitempty"`
	IsHealthy   *bool   `json:"isHealthy,omitempty"`
	LastEvent   *int64  `json:"lastEvent,omitempty"`
	TotalEvents *int64  `json:"totalEvents,omitempty"`
}

// Stakers defines model for Stakers.
type Stakers string

// StakersAddressData defines model for StakersAddressData.
type StakersAddressData struct {
	PoolsArray  *[]Asset `json:"poolsArray,omitempty"`
	TotalEarned *string  `json:"totalEarned,omitempty"`
	TotalROI    *string  `json:"totalROI,omitempty"`
	TotalStaked *string  `json:"totalStaked,omitempty"`
}

// StakersAssetData defines model for StakersAssetData.
type StakersAssetData struct {
	Asset           *Asset  `json:"asset,omitempty"`
	AssetEarned     *string `json:"assetEarned,omitempty"`
	AssetROI        *string `json:"assetROI,omitempty"`
	AssetStaked     *string `json:"assetStaked,omitempty"`
	DateFirstStaked *int64  `json:"dateFirstStaked,omitempty"`
	PoolEarned      *string `json:"poolEarned,omitempty"`
	PoolROI         *string `json:"poolROI,omitempty"`
	PoolStaked      *string `json:"poolStaked,omitempty"`
	RuneEarned      *string `json:"runeEarned,omitempty"`
	RuneROI         *string `json:"runeROI,omitempty"`
	RuneStaked      *string `json:"runeStaked,omitempty"`
	StakeUnits      *string `json:"stakeUnits,omitempty"`
}

// StatsData defines model for StatsData.
type StatsData struct {
	DailyActiveUsers   *string `json:"dailyActiveUsers,omitempty"`
	DailyTx            *string `json:"dailyTx,omitempty"`
	MonthlyActiveUsers *string `json:"monthlyActiveUsers,omitempty"`
	MonthlyTx          *string `json:"monthlyTx,omitempty"`
	PoolCount          *string `json:"poolCount,omitempty"`
	TotalAssetBuys     *string `json:"totalAssetBuys,omitempty"`
	TotalAssetSells    *string `json:"totalAssetSells,omitempty"`
	TotalDepth         *string `json:"totalDepth,omitempty"`
	TotalEarned        *string `json:"totalEarned,omitempty"`
	TotalStakeTx       *string `json:"totalStakeTx,omitempty"`
	TotalStaked        *string `json:"totalStaked,omitempty"`
	TotalTx            *string `json:"totalTx,omitempty"`
	TotalUsers         *string `json:"totalUsers,omitempty"`
	TotalVolume        *string `json:"totalVolume,omitempty"`
	TotalVolume24hr    *string `json:"totalVolume24hr,omitempty"`
	TotalWithdrawTx    *string `json:"totalWithdrawTx,omitempty"`
}

// ThorchainEndpoint defines model for ThorchainEndpoint.
type ThorchainEndpoint struct {
	Address *string `json:"address,omitempty"`
	Chain   *string `json:"chain,omitempty"`
	PubKey  *string `json:"pub_key,omitempty"`
}

// ThorchainEndpoints defines model for ThorchainEndpoints.
type ThorchainEndpoints struct {
	Current *[]ThorchainEndpoint `json:"current,omitempty"`
}

// TxDetails defines model for TxDetails.
type TxDetails struct {
	Date    *int64  `json:"date,omitempty"`
	Events  *Event  `json:"events,omitempty"`
	Gas     *Gas    `json:"gas,omitempty"`
	Height  *string `json:"height,omitempty"`
	In      *Tx     `json:"in,omitempty"`
	Options *Option `json:"options,omitempty"`
	Out     *[]Tx   `json:"out,omitempty"`
	Pool    *Asset  `json:"pool,omitempty"`
	Status  *string `json:"status,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// Asset defines model for asset.
type Asset string

// Coin defines model for coin.
type Coin struct {
	Amount *string `json:"amount,omitempty"`
	Asset  *Asset  `json:"asset,omitempty"`
}

// Coins defines model for coins.
type Coins []Coin

// Event defines model for event.
type Event struct {
	Fee        *string `json:"fee,omitempty"`
	Slip       *string `json:"slip,omitempty"`
	StakeUnits *string `json:"stakeUnits,omitempty"`
}

// Gas defines model for gas.
type Gas struct {
	Amount *string `json:"amount,omitempty"`
	Asset  *Asset  `json:"asset,omitempty"`
}

// Option defines model for option.
type Option struct {
	Asymmetry           *string `json:"asymmetry,omitempty"`
	PriceTarget         *string `json:"priceTarget,omitempty"`
	WithdrawBasisPoints *string `json:"withdrawBasisPoints,omitempty"`
}

// Tx defines model for tx.
type Tx struct {
	Address *string `json:"address,omitempty"`
	Coins   *Coins  `json:"coins,omitempty"`
	Memo    *string `json:"memo,omitempty"`
	TxID    *string `json:"txID,omitempty"`
}

// AssetsDetailedResponse defines model for AssetsDetailedResponse.
type AssetsDetailedResponse []AssetDetail

// GeneralErrorResponse defines model for GeneralErrorResponse.
type GeneralErrorResponse Error

// HealthResponse defines model for HealthResponse.
type HealthResponse struct {
	Database *bool            `json:"database,omitempty"`
	Scanners *[]ScannerStatus `json:"scanners,omitempty"`
}

// NetworkResponse defines model for NetworkResponse.
type NetworkResponse NetworkInfo

// NodeKeyResponse defines model for NodeKeyResponse.
type NodeKeyResponse []NodeKey

// PoolsDetailedResponse defines model for PoolsDetailedResponse.
type PoolsDetailedResponse []PoolDetail

// PoolsResponse defines model for PoolsResponse.
type PoolsResponse []Asset

// StakersAddressDataResponse defines model for StakersAddressDataResponse.
type StakersAddressDataResponse StakersAddressData

// StakersAssetDataResponse defines model for StakersAssetDataResponse.
type StakersAssetDataResponse []StakersAssetData

// StakersResponse defines model for StakersResponse.
type StakersResponse []Stakers

// StatsResponse defines model for StatsResponse.
type StatsResponse StatsData

// ThorchainEndpointsResponse defines model for ThorchainEndpointsResponse.
type ThorchainEndpointsResponse ThorchainEndpoints

// TxsResponse defines model for TxsResponse.
type TxsResponse struct {
	Count *int64       `json:"count,omitempty"`
	Txs   *[]TxDetails `json:"txs,omitempty"`
}

// GetAssetInfoParams defines parameters for GetAssetInfo.
type GetAssetInfoParams struct {
	Asset string `json:"asset"`
}

// GetPoolsDataParams defines parameters for GetPoolsData.
type GetPoolsDataParams struct {
	Asset string `json:"asset"`
}

// GetStakersAddressAndAssetDataParams defines parameters for GetStakersAddressAndAssetData.
type GetStakersAddressAndAssetDataParams struct {
	Asset string `json:"asset"`
}

// GetTxDetailsParams defines parameters for GetTxDetails.
type GetTxDetailsParams struct {
	Address *string `json:"address,omitempty"`
	Txid    *string `json:"txid,omitempty"`
	Asset   *string `json:"asset,omitempty"`
	Type    *string `json:"type,omitempty"`
	Offset  int64   `json:"offset"`
	Limit   int64   `json:"limit"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Asset Information// (GET /v1/assets)
	GetAssetInfo(ctx echo.Context, params GetAssetInfoParams) error
	// Get Documents// (GET /v1/doc)
	GetDocs(ctx echo.Context) error
	// Get Health// (GET /v1/health)
	GetHealth(ctx echo.Context) error
	// Get Network Data// (GET /v1/network)
	GetNetworkData(ctx echo.Context) error
	// Get Node public keys// (GET /v1/nodes)
	GetNodes(ctx echo.Context) error
	// Get Asset Pools// (GET /v1/pools)
	GetPools(ctx echo.Context) error
	// Get Pools Data// (GET /v1/pools/detail)
	GetPoolsData(ctx echo.Context, params GetPoolsDataParams) error
	// Get Stakers// (GET /v1/stakers)
	GetStakersData(ctx echo.Context) error
	// Get Staker Data// (GET /v1/stakers/{address})
	GetStakersAddressData(ctx echo.Context, address string) error
	// Get Staker Pool Data// (GET /v1/stakers/{address}/pools)
	GetStakersAddressAndAssetData(ctx echo.Context, address string, params GetStakersAddressAndAssetDataParams) error
	// Get Global Stats// (GET /v1/stats)
	GetStats(ctx echo.Context) error
	// Get Swagger// (GET /v1/swagger.json)
	GetSwagger(ctx echo.Context) error
	// Get the Proxied Pool Addresses// (GET /v1/thorchain/pool_addresses)
	GetThorchainProxiedEndpoints(ctx echo.Context) error
	// Get details of a tx by address, asset or tx-id// (GET /v1/txs)
	GetTxDetails(ctx echo.Context, params GetTxDetailsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAssetInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetInfo(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAssetInfoParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetInfo(ctx, params)
	return err
}

// GetDocs converts echo context to params.
func (w *ServerInterfaceWrapper) GetDocs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDocs(ctx)
	return err
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHealth(ctx)
	return err
}

// GetNetworkData converts echo context to params.
func (w *ServerInterfaceWrapper) GetNetworkData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNetworkData(ctx)
	return err
}

// GetNodes converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodes(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodes(ctx)
	return err
}

// GetPools converts echo context to params.
func (w *ServerInterfaceWrapper) GetPools(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPools(ctx)
	return err
}

// GetPoolsData converts echo context to params.
func (w *ServerInterfaceWrapper) GetPoolsData(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPoolsDataParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPoolsData(ctx, params)
	return err
}

// GetStakersData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersData(ctx)
	return err
}

// GetStakersAddressData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersAddressData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersAddressData(ctx, address)
	return err
}

// GetStakersAddressAndAssetData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersAddressAndAssetData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetStakersAddressAndAssetDataParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersAddressAndAssetData(ctx, address, params)
	return err
}

// GetStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetStats(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStats(ctx)
	return err
}

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetThorchainProxiedEndpoints converts echo context to params.
func (w *ServerInterfaceWrapper) GetThorchainProxiedEndpoints(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetThorchainProxiedEndpoints(ctx)
	return err
}

// GetTxDetails converts echo context to params.
func (w *ServerInterfaceWrapper) GetTxDetails(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTxDetailsParams
	// ------------- Optional query parameter "address" -------------
	if paramValue := ctx.QueryParam("address"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "address", ctx.QueryParams(), &params.Address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// ------------- Optional query parameter "txid" -------------
	if paramValue := ctx.QueryParam("txid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "txid", ctx.QueryParams(), &params.Txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	// ------------- Optional query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// ------------- Optional query parameter "type" -------------
	if paramValue := ctx.QueryParam("type"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "type", ctx.QueryParams(), &params.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// ------------- Required query parameter "offset" -------------
	if paramValue := ctx.QueryParam("offset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument offset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Required query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument limit is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTxDetails(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v1/assets", wrapper.GetAssetInfo)
	router.GET("/v1/doc", wrapper.GetDocs)
	router.GET("/v1/health", wrapper.GetHealth)
	router.GET("/v1/network", wrapper.GetNetworkData)
	router.GET("/v1/nodes", wrapper.GetNodes)
	router.GET("/v1/pools", wrapper.GetPools)
	router.GET("/v1/pools/detail", wrapper.GetPoolsData)
	router.GET("/v1/stakers", wrapper.GetStakersData)
	router.GET("/v1/stakers/:address", wrapper.GetStakersAddressData)
	router.GET("/v1/stakers/:address/pools", wrapper.GetStakersAddressAndAssetData)
	router.GET("/v1/stats", wrapper.GetStats)
	router.GET("/v1/swagger.json", wrapper.GetSwagger)
	router.GET("/v1/thorchain/pool_addresses", wrapper.GetThorchainProxiedEndpoints)
	router.GET("/v1/txs", wrapper.GetTxDetails)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"SgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

