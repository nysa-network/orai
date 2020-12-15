package provider

import (
	"github.com/oraichain/orai/x/provider/keeper"
	"github.com/oraichain/orai/x/provider/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	ScriptPath        = types.ScriptPath
)

var (
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	NewMsgCreateOracleScript = types.NewMsgCreateOracleScript
	NewMsgCreateAIDataSource = types.NewMsgCreateAIDataSource
	//NewMsgSetAIRequest       = types.NewMsgSetAIRequest
	ModuleCdc                  = types.ModuleCdc
	RegisterCodec              = types.RegisterCodec
	NewGenesisState            = types.NewGenesisState
	OScriptKeyPrefix           = types.OScriptKeyPrefix
	DataSourceKeyPrefix        = types.DataSourceKeyPrefix
	TestCaseKeyPrefix          = types.TestCaseKeyPrefix
	DataSourceStoreKeyString   = types.DataSourceStoreKeyString
	TestCaseStoreKeyString     = types.TestCaseStoreKeyString
	OracleScriptStoreKeyString = types.OracleScriptStoreKeyString
	DataSourceStoreFileString  = types.DataSourceStoreFileString
	TestCaseStoreFileString    = types.TestCaseStoreFileString
)

type (
	Keeper                = keeper.Keeper
	MsgCreateOracleScript = types.MsgCreateOracleScript
	MsgEditOracleScript   = types.MsgEditOracleScript
	MsgCreateAIDataSource = types.MsgCreateAIDataSource
	MsgEditAIDataSource   = types.MsgEditAIDataSource
	MsgCreateTestCase     = types.MsgCreateTestCase
	MsgEditTestCase       = types.MsgEditTestCase
	//MsgSetAIRequest       = types.MsgSetAIRequest
	QueryResOracleScript = types.QueryResOracleScript
	QueryResAIDataSource = types.QueryResAIDataSource
	OracleScript         = types.OracleScript
	AIDataSource         = types.AIDataSource
	//AIRequest             = types.AIRequest
	TestCase     = types.TestCase
	GenesisState = types.GenesisState
)
