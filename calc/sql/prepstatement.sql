# Copyright 2015, Google, Inc.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
SELECT 		r.airline, 
			r.source_code, a_s.name as source_name, a_s.lat as source_lat, a_s.lon as source_lon, 
			r.dest_code, a_d.name as dest_name, a_d.lat as dest_lat, a_d.lon as dest_lon
FROM 		route r
INNER JOIN 	airport a_s on r.source_id=a_s.id
INNER JOIN 	airport a_d on r.dest_id= a_d.id
