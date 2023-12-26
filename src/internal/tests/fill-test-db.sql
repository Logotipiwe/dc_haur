truncate table questions;
truncate table decks;

INSERT INTO decks
values ('d1', 'deck d1 name', 'Deck 1 desc');
INSERT INTO decks
values ('d2', 'deck d2 name', 'Deck 2 desc');
INSERT INTO decks
values ('d3', 'deck d3 name', 'Deck 3 desc');

INSERT INTO questions
values ('d1l1q1', 'l1', 'd1', 'question d1l1q1 text'),
       ('d1l1q2', 'l1', 'd1', 'question d1l1q2 text'),
       ('d1l1q3', 'l1', 'd1', 'question d1l1q3 text'),
       ('d1l2q1', 'l2', 'd1', 'question d1l2q1 text'),
       ('d1l2q2', 'l2', 'd1', 'question d1l2q2 text'),
       ('d1l3q1', 'l3', 'd1', 'question d1l3q1 text'),
       ('d1l3q2', 'l3', 'd1', 'question d1l3q2 text'),
       ('d1l3q3', 'l3', 'd1', 'question d1l3q2 text'),
       ('d2l1q1', 'l1', 'd2', 'question d2l1q1 text'),
       ('d2l2q1', 'l2', 'd2', 'question d2l2q1 text'),
       ('d2l3q1', 'l3', 'd2', 'question d2l3q1 text'),
       ('d3l1q1', 'l1', 'd3', 'question d3l1q1 text'),
       ('d3l1q2', 'l1', 'd3', 'question d3l1q2 text'),
       ('d3l2q1', 'l2', 'd3', 'question d3l2q1 text');