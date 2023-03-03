package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// Context
type Context struct{}

// -------- Abstract Layout -------

// Component interface
type Component interface {
	// Mount a component
	Mount(c Component, components ...Component) error
	// Remove a component
	Remove(c Component) error
	// Execute component & child component
	Do(ctx *Context) error
}

// Base Component
type BaseComponent struct {
	ChildComponents []Component
}

// Mount mount a component or multiples
func (bc *BaseComponent) Mount(c Component, components ...Component) (err error) {
	bc.ChildComponents = append(bc.ChildComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildComponents = append(bc.ChildComponents, components...)
	return
}

// Remove remove a component
func (bc *BaseComponent) Remove(c Component) (err error) {
	if len(bc.ChildComponents) == 0 {
		return
	}
	for k, childComponent := range bc.ChildComponents {
		if childComponent == c {
			fmt.Println(runFuncName(), " Remove:", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	return
}

// Do executing Do logic
func (bc *BaseComponent) Do(ctx *Context) (err error) {
	// Do nothing
	return
}

// ChildsDo execute childs logic
func (bc *BaseComponent) ChildsDo(ctx *Context) (err error) {
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx); err != nil {
			return err
		}
	}
	return
}

// -------- Implement Layout  ---------

// CheckOutPageComponent
type CheckOutPageComponent struct {
	BaseComponent
}

func (bc *CheckOutPageComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Order checkout page component")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// AddressComponent
type AddressComponent struct {
	BaseComponent
}

func (bc *AddressComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Address component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// Pay method component
type PayMethodComponent struct {
	BaseComponent
}

func (bc *PayMethodComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Pay method component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// Store component
type StoreComponent struct {
	BaseComponent
}

func (bc *StoreComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Store component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// SkuComponent sku goods component
type SkuComponent struct {
	BaseComponent
}

func (bc *SkuComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Sku component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// PromotionComponent
type PromotionComponent struct {
	BaseComponent
}

func (bc *PromotionComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Promotion component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// ExpressComponet
type ExpressComponent struct {
	BaseComponent
}

func (bc *ExpressComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Express component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// AfterSaleComponet
type AfterSaleComponent struct {
	BaseComponent
}

func (bc *AfterSaleComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " After Sale component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// InvoiceComponent
type InvoiceComponent struct {
	BaseComponent
}

func (bc *InvoiceComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Invoice component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// CouponComponet
type CouponComponent struct {
	BaseComponent
}

func (bc *CouponComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Couopon component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// GiftCardComponent
type GiftCardComponet struct {
	BaseComponent
}

func (bc *GiftCardComponet) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Gift component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

// OrderComponent
type OrderComponent struct {
	BaseComponent
}

func (bc *OrderComponent) Do(ctx *Context) (err error) {
	// Current component logic
	fmt.Println(runFuncName(), " Order component...")

	// Execute child component
	bc.ChildsDo(ctx)

	// Current component logic

	return
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {

	// Initlization Order page Component
	checkoutPage := &CheckOutPageComponent{}

	// Mount child components
	storeComponent := &StoreComponent{}
	skuComponent := &SkuComponent{}
	skuComponent.Mount(
		&PromotionComponent{},
		&AfterSaleComponent{},
	)
	storeComponent.Mount(
		skuComponent,
		&ExpressComponent{},
	)

	// Mount Components
	checkoutPage.Mount(
		&AddressComponent{},
		&PayMethodComponent{},
		storeComponent,
		&InvoiceComponent{},
		&CouponComponent{},
		&GiftCardComponet{},
		&OrderComponent{},
	)

	// Remove component test
	// checkoutPage.Remove(storeComponent)

	// Build page component data
	checkoutPage.Do(&Context{})

	fmt.Println("Main func done!")
}
