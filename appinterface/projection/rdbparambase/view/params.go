package view

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbparambase/types"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

// A "Params" compatible table should have the follow table schema
// | Column | Type    | Constraint  |
// | ------ | ------- | ----------- |
// | module | VARCHAR | PRIMARY KEY |
// | key    | VARCHAR | PRIMARY KEY |
// | value  | VARCHAR | NOT NULL    |

type Params interface {
	Set(accessor types.ParamAccessor, value string) error
	FindBy(accessor types.ParamAccessor) (string, error)
	FindInt64By(accessor types.ParamAccessor) (int64, error)
	FindBigIntBy(accessor types.ParamAccessor) (*big.Int, error)
	FindUint64By(accessor types.ParamAccessor) (uint64, error)
	FindDurationBy(accessor types.ParamAccessor) (time.Duration, error)
	FindBoolBy(accessor types.ParamAccessor) (bool, error)
}

// A generic param view
type ParamsView struct {
	rdbHandle *rdb.Handle

	tableName string
}

func NewParamsView(rdbHandle *rdb.Handle, tableName string) Params {
	return &ParamsView{
		rdbHandle,

		tableName,
	}
}

func (view *ParamsView) Set(accessor types.ParamAccessor, value string) error {
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.
		Insert(view.tableName).
		Columns("module", "key", "value").
		Values(accessor.Module, accessor.Key, value).
		Suffix("ON CONFLICT (module, key) DO UPDATE SET value = EXCLUDED.value").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building param insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdbHandle.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting param: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error updating param: no row updated: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *ParamsView) FindBy(accessor types.ParamAccessor) (string, error) {
	sql, sqlArgs, err := view.rdbHandle.StmtBuilder.Select(
		"value",
	).From(
		view.tableName,
	).Where(
		"module = ?", accessor.Module,
	).Where(
		"key = ?", accessor.Key,
	).ToSql()
	if err != nil {
		return "", fmt.Errorf("error preparing param selection SQL: %v", err)
	}

	var value string
	if err := view.rdbHandle.QueryRow(sql, sqlArgs...).Scan(&value); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("error getting param: %v", err)
	}

	return value, nil
}

func (view *ParamsView) FindInt64By(accessor types.ParamAccessor) (int64, error) {
	value, err := view.FindBy(accessor)
	if err != nil {
		return int64(0), err
	}

	int64Val, parseErr := strconv.ParseInt(value, 10, 64)
	if parseErr != nil {
		return int64(0), fmt.Errorf("error parsing uint64 param: %v", parseErr)
	}

	return int64Val, nil
}

func (view *ParamsView) FindBigIntBy(accessor types.ParamAccessor) (*big.Int, error) {
	value, err := view.FindBy(accessor)
	if err != nil {
		return nil, err
	}

	bigIntVal, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, errors.New("error parsing big.Int param: invalid big.Int")
	}

	return bigIntVal, nil
}

func (view *ParamsView) FindUint64By(accessor types.ParamAccessor) (uint64, error) {
	value, err := view.FindBy(accessor)
	if err != nil {
		return uint64(0), err
	}

	uint64Val, parseErr := strconv.ParseUint(value, 10, 64)
	if parseErr != nil {
		return uint64(0), fmt.Errorf("error parsing uint64 param: %v", parseErr)
	}

	return uint64Val, nil
}

func (view *ParamsView) FindDurationBy(accessor types.ParamAccessor) (time.Duration, error) {
	value, err := view.FindBy(accessor)
	if err != nil {
		return 0 * time.Second, err
	}

	durationVal, parseErr := time.ParseDuration(value)
	if parseErr != nil {
		return 0 * time.Second, fmt.Errorf("error parsing time.Duration param: %v", parseErr)
	}

	return durationVal, nil
}

func (view *ParamsView) FindBoolBy(accessor types.ParamAccessor) (bool, error) {
	value, err := view.FindBy(accessor)
	if err != nil {
		return false, err
	}

	lowerVal := strings.ToLower(value)
	if lowerVal == "true" {
		return true, nil
	} else if lowerVal == "false" {
		return false, nil
	}

	return false, errors.New("error parsing bool param: invalid boolean")
}
