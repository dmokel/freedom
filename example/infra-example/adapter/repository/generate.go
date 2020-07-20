// Code generated by 'freedom new-po'
package repository

import (
	"github.com/8treenet/freedom"
	"github.com/8treenet/freedom/example/infra-example/domain/po"
	"github.com/jinzhu/gorm"
	"time"
)

func ormErrorLog(repo freedom.GORMRepository, model, method string, e error, expression ...interface{}) {
	if e == nil || e == gorm.ErrRecordNotFound {
		return
	}
	repo.GetWorker().Logger().Errorf("Orm error, model: %s, method: %s, expression :%v, reason for error:%v", model, method, expression, e)
}

// findCart .
func findCart(repo freedom.GORMRepository, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCart", e, now)
		ormErrorLog(repo, "Cart", "findCart", e, result)
	}()
	db := repo.DB()
	if len(builders) == 0 {
		e = db.Where(result).Last(result).Error
		return
	}
	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findCartListByPrimarys .
func findCartListByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Cart", "findCartListByPrimarys", e, now)
	ormErrorLog(repo, "Cart", "findCartsByPrimarys", e, primarys)
	return
}

// findCartByWhere .
func findCartByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartByWhere", e, now)
		ormErrorLog(repo, "Cart", "findCartByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findCartByMap .
func findCartByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartByMap", e, now)
		ormErrorLog(repo, "Cart", "findCartByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findCartList .
func findCartList(repo freedom.GORMRepository, query po.Cart, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartList", e, now)
		ormErrorLog(repo, "Cart", "findCarts", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findCartListByWhere .
func findCartListByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartListByWhere", e, now)
		ormErrorLog(repo, "Cart", "findCartsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findCartListByMap .
func findCartListByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Cart", "findCartListByMap", e, now)
		ormErrorLog(repo, "Cart", "findCartsByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createCart .
func createCart(repo freedom.GORMRepository, object *po.Cart) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Cart", "createCart", e, now)
	ormErrorLog(repo, "Cart", "createCart", e, *object)
	return
}

// saveCart .
func saveCart(repo freedom.GORMRepository, object *po.Cart) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Cart", "saveCart", e, now)
	ormErrorLog(repo, "Cart", "saveCart", e, *object)
	return
}

// findDelivery .
func findDelivery(repo freedom.GORMRepository, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDelivery", e, now)
		ormErrorLog(repo, "Delivery", "findDelivery", e, result)
	}()
	db := repo.DB()
	if len(builders) == 0 {
		e = db.Where(result).Last(result).Error
		return
	}
	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findDeliveryListByPrimarys .
func findDeliveryListByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryListByPrimarys", e, now)
	ormErrorLog(repo, "Delivery", "findDeliverysByPrimarys", e, primarys)
	return
}

// findDeliveryByWhere .
func findDeliveryByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryByWhere", e, now)
		ormErrorLog(repo, "Delivery", "findDeliveryByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findDeliveryByMap .
func findDeliveryByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryByMap", e, now)
		ormErrorLog(repo, "Delivery", "findDeliveryByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findDeliveryList .
func findDeliveryList(repo freedom.GORMRepository, query po.Delivery, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryList", e, now)
		ormErrorLog(repo, "Delivery", "findDeliverys", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findDeliveryListByWhere .
func findDeliveryListByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryListByWhere", e, now)
		ormErrorLog(repo, "Delivery", "findDeliverysByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findDeliveryListByMap .
func findDeliveryListByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Delivery", "findDeliveryListByMap", e, now)
		ormErrorLog(repo, "Delivery", "findDeliverysByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createDelivery .
func createDelivery(repo freedom.GORMRepository, object *po.Delivery) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Delivery", "createDelivery", e, now)
	ormErrorLog(repo, "Delivery", "createDelivery", e, *object)
	return
}

// saveDelivery .
func saveDelivery(repo freedom.GORMRepository, object *po.Delivery) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Delivery", "saveDelivery", e, now)
	ormErrorLog(repo, "Delivery", "saveDelivery", e, *object)
	return
}

// findGoods .
func findGoods(repo freedom.GORMRepository, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoods", e, now)
		ormErrorLog(repo, "Goods", "findGoods", e, result)
	}()
	db := repo.DB()
	if len(builders) == 0 {
		e = db.Where(result).Last(result).Error
		return
	}
	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findGoodsListByPrimarys .
func findGoodsListByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsListByPrimarys", e, now)
	ormErrorLog(repo, "Goods", "findGoodssByPrimarys", e, primarys)
	return
}

// findGoodsByWhere .
func findGoodsByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsByWhere", e, now)
		ormErrorLog(repo, "Goods", "findGoodsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findGoodsByMap .
func findGoodsByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsByMap", e, now)
		ormErrorLog(repo, "Goods", "findGoodsByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findGoodsList .
func findGoodsList(repo freedom.GORMRepository, query po.Goods, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsList", e, now)
		ormErrorLog(repo, "Goods", "findGoodss", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findGoodsListByWhere .
func findGoodsListByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsListByWhere", e, now)
		ormErrorLog(repo, "Goods", "findGoodssByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findGoodsListByMap .
func findGoodsListByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Goods", "findGoodsListByMap", e, now)
		ormErrorLog(repo, "Goods", "findGoodssByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createGoods .
func createGoods(repo freedom.GORMRepository, object *po.Goods) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Goods", "createGoods", e, now)
	ormErrorLog(repo, "Goods", "createGoods", e, *object)
	return
}

// saveGoods .
func saveGoods(repo freedom.GORMRepository, object *po.Goods) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Goods", "saveGoods", e, now)
	ormErrorLog(repo, "Goods", "saveGoods", e, *object)
	return
}

// findOrder .
func findOrder(repo freedom.GORMRepository, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrder", e, now)
		ormErrorLog(repo, "Order", "findOrder", e, result)
	}()
	db := repo.DB()
	if len(builders) == 0 {
		e = db.Where(result).Last(result).Error
		return
	}
	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderListByPrimarys .
func findOrderListByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Order", "findOrderListByPrimarys", e, now)
	ormErrorLog(repo, "Order", "findOrdersByPrimarys", e, primarys)
	return
}

// findOrderByWhere .
func findOrderByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderByWhere", e, now)
		ormErrorLog(repo, "Order", "findOrderByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderByMap .
func findOrderByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderByMap", e, now)
		ormErrorLog(repo, "Order", "findOrderByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findOrderList .
func findOrderList(repo freedom.GORMRepository, query po.Order, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderList", e, now)
		ormErrorLog(repo, "Order", "findOrders", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findOrderListByWhere .
func findOrderListByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderListByWhere", e, now)
		ormErrorLog(repo, "Order", "findOrdersByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findOrderListByMap .
func findOrderListByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Order", "findOrderListByMap", e, now)
		ormErrorLog(repo, "Order", "findOrdersByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createOrder .
func createOrder(repo freedom.GORMRepository, object *po.Order) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Order", "createOrder", e, now)
	ormErrorLog(repo, "Order", "createOrder", e, *object)
	return
}

// saveOrder .
func saveOrder(repo freedom.GORMRepository, object *po.Order) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Order", "saveOrder", e, now)
	ormErrorLog(repo, "Order", "saveOrder", e, *object)
	return
}

// findAdmin .
func findAdmin(repo freedom.GORMRepository, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdmin", e, now)
		ormErrorLog(repo, "Admin", "findAdmin", e, result)
	}()
	db := repo.DB()
	if len(builders) == 0 {
		e = db.Where(result).Last(result).Error
		return
	}
	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAdminListByPrimarys .
func findAdminListByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
	now := time.Now()
	e = repo.DB().Find(results, primarys).Error
	freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminListByPrimarys", e, now)
	ormErrorLog(repo, "Admin", "findAdminsByPrimarys", e, primarys)
	return
}

// findAdminByWhere .
func findAdminByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminByWhere", e, now)
		ormErrorLog(repo, "Admin", "findAdminByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAdminByMap .
func findAdminByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminByMap", e, now)
		ormErrorLog(repo, "Admin", "findAdminByMap", e, query)
	}()

	db := repo.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAdminList .
func findAdminList(repo freedom.GORMRepository, query po.Admin, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminList", e, now)
		ormErrorLog(repo, "Admin", "findAdmins", e, query)
	}()
	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findAdminListByWhere .
func findAdminListByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminListByWhere", e, now)
		ormErrorLog(repo, "Admin", "findAdminsByWhere", e, query, args)
	}()
	db := repo.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// findAdminListByMap .
func findAdminListByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Admin", "findAdminListByMap", e, now)
		ormErrorLog(repo, "Admin", "findAdminsByMap", e, query)
	}()

	db := repo.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(results).Error
		return
	}
	e = builders[0].Execute(db, results)
	return
}

// createAdmin .
func createAdmin(repo freedom.GORMRepository, object *po.Admin) (rowsAffected int64, e error) {
	now := time.Now()
	db := repo.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	freedom.Prometheus().OrmWithLabelValues("Admin", "createAdmin", e, now)
	ormErrorLog(repo, "Admin", "createAdmin", e, *object)
	return
}

// saveAdmin .
func saveAdmin(repo freedom.GORMRepository, object *po.Admin) (affected int64, e error) {
	now := time.Now()
	db := repo.DB().Model(object).Updates(object.TakeChanges())
	e = db.Error
	affected = db.RowsAffected
	freedom.Prometheus().OrmWithLabelValues("Admin", "saveAdmin", e, now)
	ormErrorLog(repo, "Admin", "saveAdmin", e, *object)
	return
}
