insert into articles (title, contents, username, nice, created_at) values 
  ('firstPost', 'This is my first blog', 'oose', 2, now());

insert into articles (title, contents, username, nice) values 
  ('secondPost', 'This is my second blog', 'oose', 4);

insert into comments (article_id, message, created_at) values 
  (1, 'This is my first comment', now());

insert into comments (article_id, message) values 
  (1, 'welcome');