require 'mysql'
require 'fileutils'  
require "benchmark"


def cleanDir(path_to_clean)
	FileUtils.rm_rf(path_to_clean) 
	Dir.mkdir(path_to_clean,0777)
end

def deg2rad(deg)  
	deg * Math::PI / 180
end

def getDistance(lat1, lon1, lat2, lon2) 
	earth_radius = 3963

	dLat = deg2rad(lat2 - lat1)
	dLon = deg2rad(lon2 - lon1)

	a = Math.sin(dLat/2)*Math.sin(dLat/2) + Math.cos(deg2rad(lat1))*Math.cos(deg2rad(lat2))*Math.sin(dLon/2)*Math.sin(dLon/2)
	c = 2 * Math.asin(Math.sqrt(a))
	d = earth_radius * c
end

def processRoutes(routes)
	routes.each do |r|
		r['distance'] = getDistance(r['source_lat'].to_f,r['source_lon'].to_f,r['dest_lat'].to_f,r['dest_lon'].to_f);
	end
end



def writeRoutes(routes, store_path)
	Dir.mkdir(store_path,0777)

	routeText = []
	routeText.push("<table>"+ "\n")
	
	routeText.push( '	<tr>'+ "\n" +
		"		<th>Airline</th>" + "\n" +
		"		<th>Origin Aiport Code</th>" + "\n" +
		"		<th>Origin Aiport Name</th>" + "\n" +
		"		<th>Origin Latitude</th>" + "\n" +
		"		<th>Origin Longitude</th>" + "\n" +
		"		<th>Destination Aiport Code</th>" + "\n" +
		"		<th>Destination Aiport Name</th>" + "\n" +
		"		<th>Destination Latitude</th>" + "\n" +
		"		<th>Destination Longitude</th>" + "\n" +
		"		<th>Distance</th>" + "\n" +
		'	</tr>'+ "\n")

	routes.each do |route|
		routeText.push("	<tr>"+ "\n" +
		"		<td>" + route['airline'] + "</td>" + "\n" +
		"		<td>" + route['source_code'] + "</td>" + "\n" +
		"		<td>" + route['source_name'] + "</td>" + "\n" +
		"		<td>" + route['source_lat'] + "</td>" + "\n" +
		"		<td>" + route['source_lon'] + "</td>" + "\n" +
		"		<td>" + route['dest_code'] + "</td>" + "\n" +
		"		<td>" + route['dest_name'] + "</td>" + "\n" +
		"		<td>" + route['dest_lat'] + "</td>" + "\n" +
		"		<td>" + route['dest_lon'] + "</td>" + "\n" +
		"		<td>" + route['distance'].to_s + "</td>" + "\n" +
		"	</tr>"+ "\n")
	end

	routeText.push( "</table>"+ "\n")
	filename = store_path + "/table.html"
	File.open(filename, 'w') {|f| f.write(routeText.join("")) }
end

def getRoutes(mysql, sql)
	rs = mysql.query(sql) 
	result = []

	rs.each_hash do |p| 
		item = Hash.new
		item['airline'] = p['airline']
		item['source_code'] = p['source_code']
		item['source_name'] = p['source_name']
		item['source_lat'] = p['source_lat']
		item['source_lon'] = p['source_lon']
		item['dest_code'] = p['dest_code']
		item['dest_name'] = p['dest_name']
		item['dest_lat'] = p['dest_lat']
		item['dest_lon'] = p['dest_lon']
		result.push(item);
	end
	result  
end



db= Hash.new 
db['user'] = ENV["OF_USER"];
db['pass'] = ENV["OF_PASS"];
db['host'] = ENV["OF_HOST"];
db['name'] = ENV["OF_NAME"];

loopcount = ARGV[0].to_i
output_path = Dir.getwd() + "/calc/output/ruby";

sql = File.readlines('calc/sql/prepstatement.sql').join(" ")
con = Mysql.new(db['host'], db['user'], db['pass'], db['name'])  



sql += "\n" + "Limit 0," + loopcount.to_s + "\n";

cleanDir(output_path)

routes =  []
time = Benchmark.realtime do
	routes = getRoutes(con, sql)
end
puts Time.now.strftime("%y-%m-%d %H:%M:%S") + " getRoutes \t\t ElapsedTime in seconds: #{time}"

time = Benchmark.realtime do
	routes = processRoutes(routes)
end
puts Time.now.strftime("%y-%m-%d %H:%M:%S") + " processRoutes \t ElapsedTime in seconds: #{time}"

time = Benchmark.realtime do
	writeRoutes(routes, output_path+"/1/")
end
puts Time.now.strftime("%y-%m-%d %H:%M:%S") + " writeRoutes \t\t ElapsedTime in seconds: #{time}"










