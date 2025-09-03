DROP TRIGGER IF EXISTS update_orders_updated_at ON orders;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS idx_orders_created_at;
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_user_id;

DROP TABLE IF EXISTS orders;

DROP TYPE IF EXISTS order_status;
