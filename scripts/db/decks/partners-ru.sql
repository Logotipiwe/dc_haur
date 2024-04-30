delete from questions where level_id in (select id from levels where deck_id = 'partnersRu');
delete from levels where deck_id = 'partnersRu';
delete from decks where id = 'partnersRu';

INSERT INTO decks values ('partnersRu', 'RU', 'Для партнеров', '❤️', 'Для пар, чтобы лучше узнать друг друга','couples;good to start', 'heart');

INSERT INTO levels (id, deck_id, level_order, name, emoji, color_start, color_end, color_button) values
       ('f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'partnersRu', 1, 'Лайт', null, '242,62,182', '226,124,34', '164,119,105'),
       ('b49aef79-3d09-4899-8eaa-f2011e837bae', 'partnersRu', 2, 'Глубина', null, '181,3,36', '224,123,123', '105,123,164');

INSERT INTO questions
values ('2c9f652e-1f33-46be-8611-2abece8579c3', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Чем бы ты хотел(а) заниматься вместе с партнером?', null),
       ('cd6b2bb0-b999-4c9b-b6d9-55e5176c86b9', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Ты бы хотел(а) домашнее животное? Если да, то какое?', null),
       ('0151f3bb-6434-48bf-b2ce-a4917fea477b', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Какие традиции семьи ты бы хотел(а) перенести в жизнь с партнером?', null),
       ('b6e487e4-7395-4fed-bd50-3713c4a8b8c9', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Как ты любишь расслабляться?', null),
       ('1acbe67f-0108-492a-ab59-23b902c8b362', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'За какие сферы хотел(а) бы отвечать ты, а какие - делегировать партнеру?', null),
       ('cc331637-aa88-4227-a22d-a4426aea24fb', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Что ты ценишь в партнере?', null),
       ('7b604b8a-c4b9-49bc-adc2-88cc54501ff4', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Какой ты видишь свою старость?', null),
       ('197abb75-d435-4d05-ad8c-6c4767a1db03', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Какие твои цели в жизни?', null),
       ('756ecf3a-25b7-4c59-a864-f0015b6019f2', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Что тебе хочется делать в одиночестве?', null),
       ('5ee60a18-31b0-4129-97f7-bff7a11165c8', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Какие знаки внимания со стороны партнера тебе важны?', null),
       ('3fd2b9a0-f8bd-4783-b14f-32f27d51d1f8', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Каким и где ты видишь себя через 5 лет?', null),
       ('9127752c-4fbc-4e6c-b129-57d54eb7ee8a', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Как делить бюджет с партнером?', null),
       ('1f5ffa89-a490-4410-8d3d-c0e956e46001', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Что ты понимаешь под поддержкой со стороны партнера?', null),
       ('bddb2e92-7f09-4294-9161-cde82839b18e', 'f15a9b08-2fdc-4bd8-afe3-38e6ce935d9a', 'Как по-твоему, партнер мог бы себя вести во время ссоры, чтобы она быстрее рассосалась?', null),
       ('e9b3971f-2f50-4285-b035-91b9e6acee62', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Чего партнеру точно не стоит делать, чтобы не разозлить тебя?', null),
       ('6b6c8df1-cb92-43f8-ab69-1f2191ff1b13', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Ты хочешь детей? Если да, то когда?', null),
       ('2ead8c17-e1c3-49ca-be5c-ffac9545e45d', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Что ты считаешь изменой?', null),
       ('ec04cbd4-c49b-44fa-90e7-d5d741174767', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Есть ли что-то, что ты не сможешь простить партнеру?', null),
       ('8bf75761-9cc5-43a4-addb-394caf2bfc05', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Что ты сейчас чувствуешь?', null),
       ('6b4eae25-e2bc-49c8-8919-c35a34212257', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'В какие моменты тебе хочется побыть одному?', null),
       ('acbef3fe-8f60-4398-bde5-41ef1fa24e24', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Есть ли что-то, что ты давно хотел(а) спросить у партнера?', null),
       ('88781672-e831-412b-b3d0-a2c1ba17674d', 'b49aef79-3d09-4899-8eaa-f2011e837bae', 'Что новое ты бы хотел(а) попробовать в сексе?', null);
