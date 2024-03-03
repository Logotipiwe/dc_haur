truncate table questions_history;
truncate table questions;
truncate table levels;
truncate table decks;

INSERT INTO decks
values ('d1', 'deck d1 name', 'Deck 1 desc');
INSERT INTO decks
values ('d2', 'deck d2 name', 'Deck 2 desc');
INSERT INTO decks
values ('d3', 'deck d3 name', 'Deck 3 desc');

INSERT INTO levels (id, deck_id, level_order, name, color_start, color_end)
VALUES ('4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'd1', 1, 'l1', '0,0,0', '255,255,255'),
       ('dae6f634-8a6c-42a7-8d25-6a44e91e6e21', 'd1', 2, 'l2', '0,0,0', '255,255,255'),
       ('8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'd1', 3, 'l3', '0,0,0', '255,255,255'),

       ('de64eb23-9945-47fb-8da8-d8addac1dd47', 'd2', 1, 'l1', '0,0,0', '255,255,255'),
       ('d56a54ba-aeda-4a64-b290-87029b4802d5', 'd2', 2, 'l2', '0,0,0', '255,255,255'),
       ('5b0da819-cf98-435a-adb5-d36bf0778a9a', 'd2', 3, 'l3', '0,0,0', '255,255,255'),

       ('f34d9384-9181-4d0d-aa89-6d066cf77d44', 'd3', 1, 'l1', '0,0,0', '255,255,255'),
       ('95a7a834-8ff1-4ac2-a835-f75cf575697e', 'd3', 2, 'l2', '0,0,0', '255,255,255');

INSERT INTO questions
values ('d1l1q1', '4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'question d1l1q1 text'),
       ('d1l1q2', '4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'question d1l1q2 text'),
       ('d1l1q3', '4f84bde5-d6ad-4a2d-a2da-0553b4b281a2', 'question d1l1q3 text'),
       ('d1l2q1', 'dae6f634-8a6c-42a7-8d25-6a44e91e6e21', 'question d1l2q1 text'),
       ('d1l2q2', 'dae6f634-8a6c-42a7-8d25-6a44e91e6e21', 'question d1l2q2 text'),
       ('d1l3q1', '8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'question d1l3q1 text'),
       ('d1l3q2', '8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'question d1l3q2 text'),
       ('d1l3q3', '8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6', 'question d1l3q2 text'),
       ('d2l1q1', 'de64eb23-9945-47fb-8da8-d8addac1dd47', 'question d2l1q1 text'),
       ('d2l2q1', 'd56a54ba-aeda-4a64-b290-87029b4802d5', 'question d2l2q1 text'),
       ('d2l3q1', '5b0da819-cf98-435a-adb5-d36bf0778a9a', 'question d2l3q1 text'),
       ('d3l1q1', 'f34d9384-9181-4d0d-aa89-6d066cf77d44', 'question d3l1q1 text'),
       ('d3l1q2', 'f34d9384-9181-4d0d-aa89-6d066cf77d44', 'question d3l1q2 text'),
       ('d3l2q1', '95a7a834-8ff1-4ac2-a835-f75cf575697e', 'question d3l2q1 text');