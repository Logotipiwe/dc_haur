create table if not exists used_questions (
    question_id varchar(255) not null,
    client_id varchar(255) not null,
    PRIMARY KEY (question_id,client_id)
);
