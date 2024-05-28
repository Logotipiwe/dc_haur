alter table question_likes
    add column reaction_type varchar(255) not null;
update question_likes set reaction_type = 'LIKE';