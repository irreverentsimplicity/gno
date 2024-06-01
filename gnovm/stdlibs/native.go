// This file is autogenerated from the genstd tool (@/misc/genstd); do not edit.
// To regenerate it, run `go generate` from this directory.

package stdlibs

import (
	"reflect"

	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	libs_crypto_ed25519 "github.com/gnolang/gno/gnovm/stdlibs/crypto/ed25519"
	libs_crypto_sha256 "github.com/gnolang/gno/gnovm/stdlibs/crypto/sha256"
	libs_math "github.com/gnolang/gno/gnovm/stdlibs/math"
	libs_std "github.com/gnolang/gno/gnovm/stdlibs/std"
	libs_strconv "github.com/gnolang/gno/gnovm/stdlibs/strconv"
	libs_testing "github.com/gnolang/gno/gnovm/stdlibs/testing"
	libs_time "github.com/gnolang/gno/gnovm/stdlibs/time"
)

type nativeFunc struct {
	gnoPkg  string
	gnoFunc gno.Name
	params  []gno.FieldTypeExpr
	results []gno.FieldTypeExpr
	f       func(m *gno.Machine)
}

var nativeFuncs = [...]nativeFunc{
	{
		"crypto/ed25519",
		"verify",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
			{Name: gno.N("p1"), Type: gno.X("[]byte")},
			{Name: gno.N("p2"), Type: gno.X("[]byte")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []byte
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  []byte
				rp2 = reflect.ValueOf(&p2).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)

			r0 := libs_crypto_ed25519.X_verify(p0, p1, p2)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"crypto/sha256",
		"sum256",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[32]byte")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_crypto_sha256.X_sum256(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float32bits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("float32")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("uint32")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  float32
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float32bits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float32frombits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint32")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("float32")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint32
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float32frombits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float64bits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("float64")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("uint64")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  float64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float64bits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"math",
		"Float64frombits",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint64")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("float64")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint64
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_math.Float64frombits(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"bankerGetCoins",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]string")},
			{Name: gno.N("r1"), Type: gno.X("[]int64")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0, r1 := libs_std.X_bankerGetCoins(
				m,
				p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"std",
		"bankerSendCoins",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
			{Name: gno.N("p2"), Type: gno.X("string")},
			{Name: gno.N("p3"), Type: gno.X("[]string")},
			{Name: gno.N("p4"), Type: gno.X("[]int64")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  []string
				rp3 = reflect.ValueOf(&p3).Elem()
				p4  []int64
				rp4 = reflect.ValueOf(&p4).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 4, "")).TV, rp4)

			libs_std.X_bankerSendCoins(
				m,
				p0, p1, p2, p3, p4)
		},
	},
	{
		"std",
		"bankerTotalCoin",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := libs_std.X_bankerTotalCoin(
				m,
				p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"bankerIssueCoin",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
			{Name: gno.N("p2"), Type: gno.X("string")},
			{Name: gno.N("p3"), Type: gno.X("int64")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  int64
				rp3 = reflect.ValueOf(&p3).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)

			libs_std.X_bankerIssueCoin(
				m,
				p0, p1, p2, p3)
		},
	},
	{
		"std",
		"bankerRemoveCoin",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint8")},
			{Name: gno.N("p1"), Type: gno.X("string")},
			{Name: gno.N("p2"), Type: gno.X("string")},
			{Name: gno.N("p3"), Type: gno.X("int64")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint8
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  string
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  string
				rp2 = reflect.ValueOf(&p2).Elem()
				p3  int64
				rp3 = reflect.ValueOf(&p3).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 3, "")).TV, rp3)

			libs_std.X_bankerRemoveCoin(
				m,
				p0, p1, p2, p3)
		},
	},
	{
		"std",
		"emit",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("[]string")},
		},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  []string
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			libs_std.X_emit(
				m,
				p0, p1)
		},
	},
	{
		"std",
		"AssertOriginCall",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{},
		func(m *gno.Machine) {
			libs_std.AssertOriginCall(
				m,
			)
		},
	},
	{
		"std",
		"IsOriginCall",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			r0 := libs_std.IsOriginCall(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"GetChainID",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			r0 := libs_std.GetChainID(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"GetHeight",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			r0 := libs_std.GetHeight(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"origSend",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]string")},
			{Name: gno.N("r1"), Type: gno.X("[]int64")},
		},
		func(m *gno.Machine) {
			r0, r1 := libs_std.X_origSend(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"std",
		"origCaller",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			r0 := libs_std.X_origCaller(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"origPkgAddr",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			r0 := libs_std.X_origPkgAddr(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"callerAt",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_std.X_callerAt(
				m,
				p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"getRealm",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
			{Name: gno.N("r1"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0, r1 := libs_std.X_getRealm(
				m,
				p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"std",
		"derivePkgAddr",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_std.X_derivePkgAddr(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"encodeBech32",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
			{Name: gno.N("p1"), Type: gno.X("[20]byte")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  [20]byte
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := libs_std.X_encodeBech32(p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"std",
		"decodeBech32",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
			{Name: gno.N("r1"), Type: gno.X("[20]byte")},
			{Name: gno.N("r2"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0, r1, r2 := libs_std.X_decodeBech32(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r2).Elem(),
			))
		},
	},
	{
		"strconv",
		"Itoa",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_strconv.Itoa(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"AppendUint",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("[]byte")},
			{Name: gno.N("p1"), Type: gno.X("uint64")},
			{Name: gno.N("p2"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("[]byte")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  []byte
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  uint64
				rp1 = reflect.ValueOf(&p1).Elem()
				p2  int
				rp2 = reflect.ValueOf(&p2).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 2, "")).TV, rp2)

			r0 := libs_strconv.AppendUint(p0, p1, p2)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"Atoi",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int")},
			{Name: gno.N("r1"), Type: gno.X("error")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0, r1 := libs_strconv.Atoi(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
		},
	},
	{
		"strconv",
		"CanBackquote",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("bool")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_strconv.CanBackquote(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"FormatInt",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("int64")},
			{Name: gno.N("p1"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  int64
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  int
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := libs_strconv.FormatInt(p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"FormatUint",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("uint64")},
			{Name: gno.N("p1"), Type: gno.X("int")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  uint64
				rp0 = reflect.ValueOf(&p0).Elem()
				p1  int
				rp1 = reflect.ValueOf(&p1).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)
			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 1, "")).TV, rp1)

			r0 := libs_strconv.FormatUint(p0, p1)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"Quote",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_strconv.Quote(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"strconv",
		"QuoteToASCII",
		[]gno.FieldTypeExpr{
			{Name: gno.N("p0"), Type: gno.X("string")},
		},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("string")},
		},
		func(m *gno.Machine) {
			b := m.LastBlock()
			var (
				p0  string
				rp0 = reflect.ValueOf(&p0).Elem()
			)

			gno.Gno2GoValue(b.GetPointerTo(nil, gno.NewValuePathBlock(1, 0, "")).TV, rp0)

			r0 := libs_strconv.QuoteToASCII(p0)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"testing",
		"unixNano",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			r0 := libs_testing.X_unixNano()

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
		},
	},
	{
		"time",
		"now",
		[]gno.FieldTypeExpr{},
		[]gno.FieldTypeExpr{
			{Name: gno.N("r0"), Type: gno.X("int64")},
			{Name: gno.N("r1"), Type: gno.X("int32")},
			{Name: gno.N("r2"), Type: gno.X("int64")},
		},
		func(m *gno.Machine) {
			r0, r1, r2 := libs_time.X_now(
				m,
			)

			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r0).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r1).Elem(),
			))
			m.PushValue(gno.Go2GnoValue(
				m.Alloc,
				m.Store,
				reflect.ValueOf(&r2).Elem(),
			))
		},
	},
}
