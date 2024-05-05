create database if not exists `haur_test`;
use `haur_test`;

drop table if exists questions;
drop table if exists levels;
drop table if exists vector_images;
drop table if exists decks;
drop table if exists questions_history;
drop table if exists used_questions;
drop table if exists unlocked_decks;

create table if not exists decks
(
    id  varchar(255) not null primary key,
    language_code varchar(10) not null,
    name varchar(255) not null,
    emoji varchar(255),
    description text,
    labels varchar(255),
    vector_image_id varchar(255),
    hidden bool not null default false,
    promo varchar(255) unique
);
create table if not exists levels
(
    id          varchar(255) not null primary key,
    deck_id     varchar(255) not null references decks,
    level_order int          not null,
    name        varchar(255) not null,
    emoji        varchar(255),
    color_start varchar(255) not null,
    color_end   varchar(255) not null,
    color_button varchar(255) not null
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
    client_id       varchar(255) not null,
    question_time timestamp    not null default current_timestamp
);
create table if not exists used_questions (
    question_id varchar(255) not null,
    client_id varchar(255) not null,
    PRIMARY KEY (question_id,client_id)
);

create table unlocked_decks (
    client_id varchar(255),
    deck_id varchar(255),
    constraint unique index (client_id, deck_id)
);

INSERT INTO decks values ('d1', 'EN', 'deck d1 name', 'em1', 'Deck 1 desc', 'label1;label2', '1', 0, null);
INSERT INTO decks values ('d2', 'EN', 'deck d2 name', 'em2', 'Deck 2 desc', 'label1;label2', '1', 0, null);
INSERT INTO decks values ('d3', 'RU', 'deck d3 name', null, 'Deck 3 desc', 'label1;label2', '2', 0, null);
INSERT INTO decks values ('d4', 'RU', 'deck d4 name', null, 'Deck 4 desc', 'label1;label2', '2', 1, 'pro');

INSERT INTO vector_images values ('1', '<svg>1</svg>');
INSERT INTO vector_images values ('2', '<svg>2</svg>');

INSERT INTO levels (id, deck_id, level_order, name, emoji, color_start, color_end, color_button)
VALUES ('d1l1', 'd1', 1, 'l1', 'em1', '0,0,0', '255,255,255', '1,1,1'),
       ('d1l2', 'd1', 2, 'l2', 'em1', '0,0,0', '255,255,255', '2,2,2'),
       ('d1l3', 'd1', 3, 'l3', null, '0,0,0', '255,255,255', '3,3,3'),

       ('d2l1', 'd2', 1, 'l1', 'em2', '0,0,0', '255,255,255', '1,1,1'),
       ('d2l2', 'd2', 2, 'l2', null, '0,0,0', '255,255,255', '2,2,2'),
       ('d2l3', 'd2', 3, 'l3', 'em2', '0,0,0', '255,255,255', '3,3,3'),

       ('d3l1', 'd3', 1, 'l1', 'em3', '0,0,0', '255,255,255', '1,1,1'),
       ('d3l2', 'd3', 2, 'l2', null, '0,0,0', '255,255,255', '2,2,2'),

       ('d4l1', 'd4', 1, 'l1', null, '0,0,0', '255,255,255', '2,2,2'),
       ('d4l2', 'd4', 2, 'l2', null, '0,0,0', '255,255,255', '2,2,2');

INSERT INTO questions
values ('d1l1q1', 'd1l1', 'question d1l1q1 text', 'additional'),
       ('d1l1q2', 'd1l1', 'question d1l1q2 text', 'additional'),
       ('d1l1q3', 'd1l1', 'question d1l1q3 text', 'additional'),
       ('d1l2q1', 'd1l2', 'question d1l2q1 text', 'additional'),
       ('d1l2q2', 'd1l2', 'question d1l2q2 text', 'additional'),
       ('d1l3q1', 'd1l3', 'question d1l3q1 text', 'additional'),
       ('d1l3q2', 'd1l3', 'question d1l3q2 text', 'additional'),
       ('d1l3q3', 'd1l3', 'question d1l3q2 text', 'additional'),
       ('d2l1q1', 'd2l1', 'question d2l1q1 text', 'additional'),
       ('d2l2q1', 'd2l2', 'question d2l2q1 text', 'additional'),
       ('d2l3q1', 'd2l3', 'question d2l3q1 text', 'additional'),
       ('d3l1q1', 'd3l1', 'question d3l1q1 text', 'additional'),
       ('d3l1q2', 'd3l1', 'question d3l1q2 text', 'additional'),
       ('d3l2q1', 'd3l2', 'question d3l2q1 text', 'additional'),

       ('d4l1q1', 'd4l1', 'question d4l1q1 text', 'additional'),
       ('d4l1q2', 'd4l1', 'question d4l1q2 text', 'additional'),
       ('d4l2q1', 'd4l2', 'question d4l2q1 text', 'additional');