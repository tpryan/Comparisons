SELECT 		r.airline, 
			r.source_code, a_s.name as source_name, a_s.lat as source_lat, a_s.lon as source_lon, 
			r.dest_code, a_d.name as dest_name, a_d.lat as dest_lat, a_d.lon as dest_lon
FROM 		route r
INNER JOIN 	airport a_s on r.source_id=a_s.id
INNER JOIN 	airport a_d on r.dest_id= a_d.id
