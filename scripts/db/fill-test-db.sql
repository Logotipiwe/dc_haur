create database if not exists `haur_test`;
use `haur_test`;

drop table if exists questions;
drop table if exists levels;
drop table if exists vector_images;
drop table if exists decks;

create table if not exists decks
(
    id  varchar(255) not null primary key,
    name varchar(255) not null,
    emoji varchar(255),
    description text,
    labels varchar(255),
    vector_image_id varchar(255)
);
create table if not exists levels
(
    id          varchar(255) not null primary key,
    deck_id     varchar(255) not null references decks,
    level_order int          not null,
    name        varchar(255) not null,
    emoji        varchar(255),
    color_start varchar(255) not null,
    color_end   varchar(255) not null
);
create table if not exists questions
(
    id              varchar(255) not null primary key,
    level_id varchar(255) not null references levels,
    text            varchar(255) not null,
    additional_text varchar(255)
);
create table if not exists vector_images
(
    id varchar(255) not null primary key,
    content longtext not null
);

create table if not exists questions_history
(
    id            varchar(255) not null primary key,
    level_id varchar(255) not null,
    question_id   varchar(255) not null,
    chat_id       varchar(255) not null,
    question_time timestamp    not null default current_timestamp
);

INSERT INTO decks values ('d1', 'deck d1 name', 'em1', 'Deck 1 desc', 'label1;label2', '1');
INSERT INTO decks values ('d2', 'deck d2 name', 'em2', 'Deck 2 desc', 'label1;label2', '1');
INSERT INTO decks values ('d3', 'deck d3 name', null, 'Deck 3 desc', 'label1;label2', '2');

INSERT INTO vector_images values ('1', '<svg>1</svg>');
INSERT INTO vector_images values ('2', '<svg>2</svg>');

INSERT INTO levels (id, deck_id, level_order, name, emoji, color_start, color_end)
VALUES ('4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'd1', 1, 'l1', 'em1', '0,0,0', '255,255,255'),
       ('dae6f634-8a6c-42a7-8d25-6a44e91e6e21', 'd1', 2, 'l2', 'em1', '0,0,0', '255,255,255'),
       ('8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'd1', 3, 'l3', null, '0,0,0', '255,255,255'),

       ('de64eb23-9945-47fb-8da8-d8addac1dd47', 'd2', 1, 'l1', 'em2', '0,0,0', '255,255,255'),
       ('d56a54ba-aeda-4a64-b290-87029b4802d5', 'd2', 2, 'l2', null, '0,0,0', '255,255,255'),
       ('5b0da819-cf98-435a-adb5-d36bf0778a9a', 'd2', 3, 'l3', 'em2', '0,0,0', '255,255,255'),

       ('f34d9384-9181-4d0d-aa89-6d066cf77d44', 'd3', 1, 'l1', 'em3', '0,0,0', '255,255,255'),
       ('95a7a834-8ff1-4ac2-a835-f75cf575697e', 'd3', 2, 'l2', null, '0,0,0', '255,255,255');

INSERT INTO questions
values ('d1l1q1', '4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'question d1l1q1 text', 'additional'),
       ('d1l1q2', '4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'question d1l1q2 text', 'additional'),
       ('d1l1q3', '4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'question d1l1q3 text', 'additional'),
       ('d1l2q1', 'dae6f634-8a6c-42a7-8d25-6a44e91e6e21', 'question d1l2q1 text', 'additional'),
       ('d1l2q2', 'dae6f634-8a6c-42a7-8d25-6a44e91e6e21', 'question d1l2q2 text', 'additional'),
       ('d1l3q1', '8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'question d1l3q1 text', 'additional'),
       ('d1l3q2', '8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'question d1l3q2 text', 'additional'),
       ('d1l3q3', '8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'question d1l3q2 text', 'additional'),
       ('d2l1q1', 'de64eb23-9945-47fb-8da8-d8addac1dd47', 'question d2l1q1 text', 'additional'),
       ('d2l2q1', 'd56a54ba-aeda-4a64-b290-87029b4802d5', 'question d2l2q1 text', 'additional'),
       ('d2l3q1', '5b0da819-cf98-435a-adb5-d36bf0778a9a', 'question d2l3q1 text', 'additional'),
       ('d3l1q1', 'f34d9384-9181-4d0d-aa89-6d066cf77d44', 'question d3l1q1 text', 'additional'),
       ('d3l1q2', 'f34d9384-9181-4d0d-aa89-6d066cf77d44', 'question d3l1q2 text', 'additional'),
       ('d3l2q1', '95a7a834-8ff1-4ac2-a835-f75cf575697e', 'question d3l2q1 text', 'additional');