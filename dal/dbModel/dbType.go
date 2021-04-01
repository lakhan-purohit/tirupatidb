package dbModel

import "time"

type DB struct {
	EngineName         string       `json:"Engine"`
	EngineConfigStruct EngineConfig `json:"EngineConfig"`
}

type EngineConfig struct {
	Host        string `json:"Host"`
	User        string `json:"User"`
	Password    string `json:"Password"`
	DB          string `json:"Db"`
	MaxOpenConn int    `json:"MaxOpenConn"`
	MaxIdelConn int    `json:"MaxIdelConn"`
}

type UserInfo struct {
	ID               int64     `db:"id" json:"Id"`
	UserMobileNumber string    `db:"user_mobile_number" json:"UserMobileNumber"`
	UserName         string    `db:"user_name" json:"UserName"`
	EmailID          string    `db:"email_id" json:"EmailID"`
	FirmName         string    `db:"firm_name" json:"FirmName"`
	Password         string    `db:"password" json:"Password"`
	CreatedAt        time.Time `db:"created_at" json:"CreatedAt"`
	Blocked          bool      `db:"blocked" json:"Blocked"`
	Role             int       `db:"role" json:"Role"`
}

type Product struct {
	ID        int64          `db:"id" json:"Id"`
	Color     string         `db:"color" json:"Color"`
	Nack      string         `db:"nack" json:"Nack"`
	Size      string         `db:"size" json:"Size"`
	BrandName string         `db:"brand_name" json:"BrandName"`
	HandSlive string         `db:"hand_slive" json:"HandSlive"`
	Pocket    string         `db:"pocket" json:"Pocket"`
	ClothType string         `db:"cloth_type" json:"ClothType"`
	Role      string         `db:"role" json:"Role"`
	Photos    []ProductPhoto `json:"Photos"`
	Price     int64          `db:"price" json:"Price"`
}
type ProductPhoto struct {
	ID           int64  `db:"id" json:"Id"`
	Photo        string `db:"photo" json:"Photo"`
	PhotoThumb   string `db:"photo_thumb" json:"PhotoThumb"`
	ClothTableID int64  `db:"clothTable_id" json:"ClothTableId"`
}

type SaveProduct struct {
	ID        int64 `db:"id" json:"Id"`
	UserID    int64 `db:"user_id" json:"UserId"`
	ProductID int64 `db:"product_id" json:"ProductId"`
}

type MakeOrder struct {
	ID        int64     `db:"id" json:"Id"`
	UserID    int64     `db:"user_id" json:"UserId"`
	ProductID int64     `db:"product_id" json:"ProductId"`
	Qty       int       `db:"qty" json:"Qty"`
	CreatedAt time.Time `db:"createdAt" json:"CreateAt"`
	UpdateAt  time.Time `db:"updateAt" json:"UpdateAt"`
	Status    int       `db:"status" json:"Status"`
	Total     int64     `db:"total" json:"Total"`
	Paid      int64     `db:"paid" json:"Paid"`
	Price     int64     `db:"price" json:"Price"`
}

type OrderDetails struct {
	ID        int64     `db:"id" json:"Id"`
	User      UserInfo  `json:"User"`
	Product   Product   `json:"Product"`
	Qty       int       `db:"qty" json:"Qty"`
	CreatedAt time.Time `db:"create_at" json:"CreateAt"`
	UpdateAt  time.Time `db:"update_at" json:"UpdateAt"`
	Status    int       `db:"status" json:"Status"`
	Total     int64     `db:"total" json:"Total"`
	Paid      int64     `db:"paid" json:"Paid"`
	Price     int64     `db:"price" json:"Price"`
}

type Order struct {
	Order []MakeOrder `json:"order`
}

const (
	USERINFOTABLE string = "userinfo"
	PRODUCTLIST   string = "t_shirt_table"
	CLOUTHPHOTO   string = "cloth_photo"
	SAVEPRODUCT   string = "save_product"
	ORDERLIST     string = "order_list"
)

type AddProductInterface interface {
	AddProduct(Product) (int64, error)
	GetProductList(role int64) ([]Product, error)
	SaveImagesProduct(int64, string) (bool, error)
	DeleteImageProduct(int64) (bool, error)
	DeleteProduct(int64) (bool, error)
	UpdateProduct(Product) (bool, error)
	SaveUnSaveProduct(int64, int64) (bool, error)
}

type UserInterface interface {
	CreateUser(UserInfo) (UserInfo, error)
	LoginUser(UserInfo) (UserInfo, error)
	UserList() ([]UserInfo, error)
	UserDetailsUpdate(UserInfo) (UserInfo, error)
}

type OrderInterface interface {
	MakeOrderList(Order) (bool, error)
	GetOrder(int64) ([]OrderDetails, error)
	UpdateOrder(int64, int64, int64) (bool, error)
	ChangeOrderStatus(int64, int64) (bool, error)
	DeleteOrder(int64) (bool, error)
}

type RouterFunc interface {
	UserInterface
	AddProductInterface
	OrderInterface
}
