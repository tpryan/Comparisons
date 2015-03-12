require 'mysql'
require 'fileutils'


def cleanDir(path_to_clean)
  FileUtils.rm_rf(path_to_clean) 
  Dir.mkdir(path_to_clean,0777)
end

def cleanURL(url)
  url.gsub("blog//blog/index.php/", "").gsub("http://http://", "http://")
end


def getEntries(mysql, sql)
  rs = mysql.query(sql) 
  result = []

  rs.each_hash do |p| 
    item = Hash.new
    item['title'] = p['post_title']
    item['name'] = p['post_name']
    item['post_date'] = p['post_date']
    item['excerpt'] = p['post_excerpt']
    item['url'] = cleanURL(p['guid'])
    item['content'] = p['post_content']
    item['date'] = p['formatted_post_date']
    result.push(item);
  end
  result  
end


def writeEntries(entries, store_path)
  Dir.mkdir(store_path,0777);

  entries.each do |entry|
    item = ""
    item << "<article>\n"
    item << " <h1><a href=\"#{entry["url"]}\">#{entry["title"]}</a></h1>\n"
    item << " <time datetime=\"#{entry["post_date"]}\">#{entry["date"]}</time>\n"
    item << " <div>\n"
    item << entry['content']
    item << "\n"
    item << " </div>\n"
    item << "</article>\n"
    filename = store_path + "/" + entry['name'] + ".html";
    File.open(filename, 'w') {|f| f.write(item) }
  end

end
  
db= Hash.new 
db['user'] = ENV["DB_USER"];
db['pass'] = ENV["DB_PASS"];
db['host'] = ENV["DB_HOST"];
db['name'] = ENV["DB_NAME"];

loopcount = ARGV[0].to_i
output_path = Dir.getwd() + "/textout/output/ruby";

sql = File.readlines('textout/sql/entries.sql').join(" ")
con = Mysql.new(db['host'], db['user'], db['pass'], db['name'])  

cleanDir(output_path)

entries = getEntries(con, sql)

for i in 1..loopcount do
  writeEntries(entries, output_path + "/" +  i.to_s)
end
 
con.close 