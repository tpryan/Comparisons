SELECT
    p.post_title, 
    p.post_excerpt, 
    p.post_name, 
    p.guid, 
    p.post_date, 
    p.post_content, 
    DATE_FORMAT(p.post_date, '%M %d, %Y') as formatted_post_date
FROM 
    wp_posts p
WHERE
    p.post_status='publish' 
    AND 
    p.post_type='post'
ORDER BY 
    p.post_date DESC