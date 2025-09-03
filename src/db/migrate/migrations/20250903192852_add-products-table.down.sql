DROP TRIGGER IF EXISTS update_products_updated_at ON products;

DROP FUNCTION IF EXISTS update_products_updated_at();

DROP INDEX IF EXISTS idx_products_search;
DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_products_quantity;
DROP INDEX IF EXISTS idx_products_price;
DROP INDEX IF EXISTS idx_products_name;

DROP TABLE IF EXISTS products;
