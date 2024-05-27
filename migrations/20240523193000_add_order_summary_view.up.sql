CREATE OR REPLACE VIEW order_summary_view
AS
SELECT
    o.order_id,
    e.estimate_id,
    e.user_id,
    ei.estimate_item_id,
    ei.quantity,
    ei.merchant_id,
    mi.merchant_item_id,
    mi.name AS item_name,
    mi.category AS item_category,
    mi.price AS item_price,
    mi.image_url AS item_image_url,
    mi.created_at AS item_created_at
FROM
    (orders o JOIN estimates e
    ON
	 o.estimate_id = e.estimate_id)
    JOIN (estimate_items ei JOIN merchant_items mi
    ON
	  ei.merchant_item_id = mi.merchant_item_id)
    ON o.estimate_id = ei.estimate_id;