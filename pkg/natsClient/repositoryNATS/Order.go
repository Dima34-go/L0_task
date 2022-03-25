package repositoryNATS

import (
	todo "WB_GO_L0"
	"WB_GO_L0/pkg/storage"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	C *storage.Cache
	DB *sqlx.DB
}

func NewTodoOrderPSql(cache *storage.Cache,db *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		C: cache,
		DB: db,
	}
}
func (r *OrderRepo) AddOrderItems(Order *todo.OrderItems) error {
	tx,err:=r.DB.Begin()
	if err != nil {
		return err
	}
	psqlQ := fmt.Sprintf(`INSERT INTO %s (order_uid, track_number, entry, name, phone, zip, city, address, 
    region, email, transaction,request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total,
    custom_fee,locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, oof_shard, date_created) 
    VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, '%s', %d, %d, %d, '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s')
    RETURNING order_id;`, storage.OrdersTable,Order.OrderUid,Order.TrackNumber,Order.Entry,Order.Name,Order.Phone,Order.Zip,
    Order.City,Order.Address,Order.Region,Order.Email,Order.Transaction,Order.RequestId,Order.Currency,Order.Provider,Order.Amount,
    Order.PaymentDt,Order.Bank,Order.DeliveryCost,Order.GoodsTotal,Order.CustomFee,Order.Locale,Order.InternalSignature,Order.CustomerId,
    Order.DeliveryService,Order.ShardKey,Order.SmId,Order.OofShard,Order.DateCreated)

	row:=tx.QueryRow(psqlQ)

	if err=row.Scan(&Order.OrderId);err!=nil{
		tx.Rollback()
		return err
	}
	for _,Item:=range Order.Items{
		psqlQ = fmt.Sprintf(`INSERT INTO %s(
	chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_id)
	VALUES (%d, '%s', %d, '%s', '%s', %d, '%s', %d, %d, '%s', %d, %d);`,storage.ItemsTable,Item.ChrtId,Item.TrackNumber,Item.Price,Item.Rid,Item.Name,Item.Sale,Item.Size,Item.TotalPrice,Item.NmId,Item.Brand,Item.Status,Order.OrderId)
		_,err:=tx.Query(psqlQ)
		if err!=nil{
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	r.C.Add(Order.OrderId,*Order)
	return nil
}