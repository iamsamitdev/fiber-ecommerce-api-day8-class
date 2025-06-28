package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// AuthService interface สำหรับการจัดการการยืนยันตัวตน
type AuthService interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error)
	AdminRegister(ctx context.Context, req *entities.AdminRegisterRequest) (*entities.User, error)
	Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error)
	RefreshToken(ctx context.Context, req *entities.RefreshTokenRequest) (*entities.LoginResponse, error)
	Logout(ctx context.Context, userID uuid.UUID) error
	ChangePassword(ctx context.Context, userID uuid.UUID, req *entities.ChangePasswordRequest) error
	ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error
	ValidateToken(ctx context.Context, token string) (*entities.User, error)
}

// UserService interface สำหรับการจัดการผู้ใช้
type UserService interface {
	GetUsers(ctx context.Context, page, limit int) ([]*entities.User, *entities.PaginationResponse, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req *entities.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// CategoryService interface สำหรับการจัดการหมวดหมู่
type CategoryService interface {
	CreateCategory(ctx context.Context, req *entities.CreateCategoryRequest) (*entities.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]*entities.Category, *entities.PaginationResponse, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, req *entities.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}

// ProductService interface สำหรับการจัดการสินค้า
type ProductService interface {
	CreateProduct(ctx context.Context, req *entities.CreateProductRequest) (*entities.Product, error)
	GetProducts(ctx context.Context, page, limit int) ([]*entities.Product, *entities.PaginationResponse, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*entities.Product, error)
	GetProductsByCategory(ctx context.Context, categoryID uuid.UUID, page, limit int) ([]*entities.Product, *entities.PaginationResponse, error)
	SearchProducts(ctx context.Context, req *entities.ProductSearchRequest) ([]*entities.Product, *entities.PaginationResponse, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, req *entities.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

// CartService interface สำหรับการจัดการตะกร้าสินค้า
type CartService interface {
	GetCart(ctx context.Context, userID uuid.UUID) (*entities.Cart, error)
	AddToCart(ctx context.Context, userID uuid.UUID, req *entities.AddToCartRequest) error
	UpdateCartItem(ctx context.Context, cartItemID uuid.UUID, req *entities.UpdateCartItemRequest) error
	RemoveFromCart(ctx context.Context, cartItemID uuid.UUID) error
	ClearCart(ctx context.Context, userID uuid.UUID) error
}

// OrderService interface สำหรับการจัดการคำสั่งซื้อ
type OrderService interface {
	CreateOrder(ctx context.Context, userID uuid.UUID, req *entities.CreateOrderRequest) (*entities.Order, error)
	GetOrders(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entities.Order, *entities.PaginationResponse, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*entities.Order, error)
	CancelOrder(ctx context.Context, id uuid.UUID) error
	GetAllOrders(ctx context.Context, page, limit int) ([]*entities.Order, *entities.PaginationResponse, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, req *entities.UpdateOrderStatusRequest) error
	UpdatePaymentStatus(ctx context.Context, id uuid.UUID, req *entities.UpdatePaymentStatusRequest) error
	UpdateShippingStatus(ctx context.Context, id uuid.UUID, req *entities.UpdateShippingStatusRequest) error
}

// PaymentService interface สำหรับการจัดการการชำระเงิน
type PaymentService interface {
	CreatePayment(ctx context.Context, req *entities.CreatePaymentRequest) (*entities.Transaction, error)
	GetPaymentByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error)
	VerifyPayment(ctx context.Context, id uuid.UUID, req *entities.VerifyPaymentRequest) error
	CancelPayment(ctx context.Context, id uuid.UUID) error
}

// StatsService interface สำหรับสถิติ
type StatsService interface {
	GetSalesStats(ctx context.Context) (*entities.SalesStats, error)
	GetProductStats(ctx context.Context) (*entities.ProductStats, error)
	GetUserStats(ctx context.Context) (*entities.UserStats, error)
}
