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