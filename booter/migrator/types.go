package migrator

import (
	zsw "github.com/zhongshuwen/zswchain-go"
)

// Inject represents the `inject` struct on `migration` contract.
type Inject struct {
	Table zsw.TableName `json:"table"`
	Scope zsw.ScopeName `json:"scope"`
	Payer zsw.Name      `json:"payer"`
	Key   zsw.Name      `json:"id"`
	Data  zsw.HexBytes  `json:"data"`
}

func newInjectAct(account zsw.AccountName, table zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, key zsw.Name, data zsw.HexBytes) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("inject"),
		Authorization: []zsw.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: zsw.NewActionData(Inject{Table: table, Scope: scope, Payer: zsw.Name(payer), Key: key, Data: data}),
	}
}

// Idxi represents the `Idxi` struct on `migration` contract.
type Idxi struct {
	Table     zsw.TableName `json:"table"`
	Scope     zsw.ScopeName `json:"scope"`
	Payer     zsw.Name      `json:"payer"`
	Key       zsw.Name      `json:"id"`
	Secondary zsw.Name      `json:"secondary"`
}

func newIdxi(account zsw.AccountName, tableName zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, primKey zsw.Name, value zsw.Name) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("idxi"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      payer,
				Permission: PN("active"),
			},
		},
		ActionData: zsw.NewActionData(Idxi{
			Table:     tableName,
			Scope:     scope,
			Payer:     zsw.Name(payer),
			Key:       primKey,
			Secondary: value,
		}),
	}
}

// Idxii represents the `Idxii` struct on `migration` contract.
type Idxii struct {
	Table     zsw.TableName `json:"table"`
	Scope     zsw.ScopeName `json:"scope"`
	Payer     zsw.Name      `json:"payer"`
	Key       zsw.Name      `json:"id"`
	Secondary zsw.Uint128   `json:"secondary"`
}

func newIdxii(account zsw.AccountName, tableName zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, primKey zsw.Name, value zsw.Uint128) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("idxii"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      payer,
				Permission: PN("active"),
			},
		},
		ActionData: zsw.NewActionData(Idxii{
			Table:     tableName,
			Scope:     scope,
			Payer:     zsw.Name(payer),
			Key:       primKey,
			Secondary: value,
		}),
	}
}

// Idxc represents the `Idxc` struct on `migration` contract.
type Idxc struct {
	Table     zsw.TableName   `json:"table"`
	Scope     zsw.ScopeName   `json:"scope"`
	Payer     zsw.Name        `json:"payer"`
	Key       zsw.Name        `json:"id"`
	Secondary zsw.Checksum256 `json:"secondary"`
}

func newIdxc(account zsw.AccountName, tableName zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, primKey zsw.Name, value zsw.Checksum256) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("idxc"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      payer,
				Permission: PN("active"),
			},
		},
		ActionData: zsw.NewActionData(Idxc{
			Table:     tableName,
			Scope:     scope,
			Payer:     zsw.Name(payer),
			Key:       primKey,
			Secondary: value,
		}),
	}
}

// Idxdbl represents the `Idxdbl` struct on `migration` contract.
type Idxdbl struct {
	Table     zsw.TableName `json:"table"`
	Scope     zsw.ScopeName `json:"scope"`
	Payer     zsw.Name      `json:"payer"`
	Key       zsw.Name      `json:"id"`
	Secondary float64       `json:"secondary"`
}

func newIdxdbl(account zsw.AccountName, tableName zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, primKey zsw.Name, value float64) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("idxdbl"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      payer,
				Permission: PN("active"),
			},
		},
		ActionData: zsw.NewActionData(Idxdbl{
			Table:     tableName,
			Scope:     scope,
			Payer:     zsw.Name(payer),
			Key:       primKey,
			Secondary: value,
		}),
	}
}

// Idxldbl represents the `Idxldbl` struct on `migration` contract.
type Idxldbl struct {
	Table     zsw.TableName `json:"table"`
	Scope     zsw.ScopeName `json:"scope"`
	Payer     zsw.Name      `json:"payer"`
	Key       zsw.Name      `json:"id"`
	Secondary zsw.Float128  `json:"secondary"`
}

func newIdxldbl(account zsw.AccountName, tableName zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, primKey zsw.Name, value zsw.Float128) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("idxldbl"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      payer,
				Permission: PN("active"),
			},
		},
		ActionData: zsw.NewActionData(Idxldbl{
			Table:     tableName,
			Scope:     scope,
			Payer:     zsw.Name(payer),
			Key:       primKey,
			Secondary: value,
		}),
	}
}

// Delete represents the `Delete` struct on `migration` contract.
type Eject struct {
	Account zsw.AccountName `json:"account"`
	Table   zsw.TableName   `json:"table"`
	Scope   zsw.ScopeName   `json:"scope"`
	Key     zsw.Name        `json:"id"`
}

func newEject(account zsw.AccountName, tableName zsw.TableName, scope zsw.ScopeName, payer zsw.AccountName, primKey zsw.Name) *zsw.Action {
	return &zsw.Action{
		Account: account,
		Name:    ActN("eject"),
		Authorization: []zsw.PermissionLevel{
			{
				Actor:      payer,
				Permission: PN("active"),
			},
		},
		ActionData: zsw.NewActionData(Eject{
			Account: account,
			Table:   tableName,
			Scope:   scope,
			Key:     primKey,
		}),
	}
}
