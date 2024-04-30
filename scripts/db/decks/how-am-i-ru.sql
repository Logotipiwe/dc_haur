delete from questions where level_id in (select id from levels where deck_id = 'howAmIru');
delete from levels where deck_id = 'howAmIru';
delete from decks where id = 'howAmIru';

INSERT INTO decks values ('howAmIru', 'RU', 'Как у меня дела?', '😌', 'Колода для само-рефлексии о себе и своей жизни. Удели каждому вопросу больше времени, подумай о нём с разных сторон, и он поможет тебе открыть новые черты своей личности.', 'good to start;self-care', 'b13759e3-0582-41f1-b882-89d1296f5e3c');

INSERT INTO levels (id, deck_id, level_order, name, emoji, color_start, color_end, color_button)
    VALUES ('14b225d6-6a7c-442c-971b-a95e437a3d23', 'howAmIru', 1, 'Знакомство', null, '26,59,153', '95,150,51', '121,156,128');

INSERT INTO questions values ('972f7a07-9a3b-4337-b77d-dbee8e7aea65','14b225d6-6a7c-442c-971b-a95e437a3d23','Есть ли у меня кумир или пример для подражания? Что в нём/ней меня привлекает? Могу ли я попробовать добиться похожих успехов?',null),
                             ('6f277a88-4e30-4752-8d5b-da2d98e5c208','14b225d6-6a7c-442c-971b-a95e437a3d23','Какая сейчас главная сфера жизни, над которой я работаю? Что мне в ней ещё нужно сделать?',null),
                             ('e600e31f-e75b-4945-a5b0-fb645ac4801c','14b225d6-6a7c-442c-971b-a95e437a3d23','Книги на какую тему мне сейчас хочется прочесть? Смогу ли я выделить на это время?',null),
                             ('5dcb144b-2063-495b-a55f-ea3805c5085e','14b225d6-6a7c-442c-971b-a95e437a3d23','Сохраняю ли я баланс между работой и личной жизнью?',null),
                             ('7885a621-15e1-44da-93c6-bad539c310e6','14b225d6-6a7c-442c-971b-a95e437a3d23','Какие слова мне важнее всего слышать?',null),
                             ('5a374ad6-f337-4f4d-a164-f2ee7e763e39','14b225d6-6a7c-442c-971b-a95e437a3d23','Что для меня счастье?',null),
                             ('a8102c42-0680-45a7-904c-89e9e6e41faf','14b225d6-6a7c-442c-971b-a95e437a3d23','Составь список всего, что тебя вдохновляет: образы, книги, цитаты, места','Действие'),
                             ('ad9435e5-29bf-43a1-bcd0-47307d8ae137','14b225d6-6a7c-442c-971b-a95e437a3d23','Когда я ощущаю наибольшую энергичность? Могу ли я внести больше этого в свою жизнь?',null),
                             ('1ecaa431-edcb-44c5-90c8-4d8e854cb8bf','14b225d6-6a7c-442c-971b-a95e437a3d23','Получил ли кто-то сильную поддержку от меня в последнее время? Могу ли я оказать такую же поддержку себе?',null),
                             ('10d819e9-2a19-4200-8964-9b5c96da22ea','14b225d6-6a7c-442c-971b-a95e437a3d23','Какой мой труд был самым реальным, необходимым и счастливым? Что позволило ему случиться?',null),
                             ('ecce8342-0ef9-4a91-bf24-ac1c9c74017d','14b225d6-6a7c-442c-971b-a95e437a3d23','Есть ли для меня различие между “жить” и “существовать”? В чём оно?',null),
                             ('b5ee09b4-fb09-4a89-b484-56e147d769be','14b225d6-6a7c-442c-971b-a95e437a3d23','Что я хочу чтобы другие знали обо мне, чего сейчас не знают? Почему мне важно донести это до них?',null),
                             ('aa92935c-fb70-41cb-822e-ed648cb1fc9d','14b225d6-6a7c-442c-971b-a95e437a3d23','Забочусь ли я о себе так же, как о близких мне людях? Стоит ли мне уделять больше времени себе?',null),
                             ('a4c3f857-4550-4ef3-8d32-b4ea381a301b','14b225d6-6a7c-442c-971b-a95e437a3d23','С кем я сейчас хочу встретиться или поговорить?',null),
                             ('146272be-faa8-4e5a-9822-074e669c2328','14b225d6-6a7c-442c-971b-a95e437a3d23','Какого персонажа я больше всего асоциирую с собой? Какие качества/черты мне в нем нравятся? Хочу ли я иметь похожую на него историю?',null),
                             ('ebfa241a-de81-4d02-9ba6-2d202434ad45','14b225d6-6a7c-442c-971b-a95e437a3d23','Встретив себя на 5 лет моложе, какую фразу я скажу?',null),
                             ('bac77f1c-8095-4570-b1f9-4c241df89ebc','14b225d6-6a7c-442c-971b-a95e437a3d23','Какой мой самый любимый способ провести день?',null),
                             ('2fbe3c5e-bc52-4f87-8469-63d15c8e73bb','14b225d6-6a7c-442c-971b-a95e437a3d23','Какое для меня значение имеет мнение окружающих обо мне? Стоит ли мне придавать этому больше или меньше значения?',null),
                             ('842d0026-a165-4f0f-81cb-af3d6df2c571','14b225d6-6a7c-442c-971b-a95e437a3d23','Кто или что оказало наибольшее влияние на мою жизнь сейчас?',null),
                             ('b0f2a62b-9645-49d8-832c-983a3486919d','14b225d6-6a7c-442c-971b-a95e437a3d23','Скольким из моих друзей можно было бы доверить свою жизнь?',null),
                             ('e661feb6-7e99-4baa-a0ae-0bf81b6ab050','14b225d6-6a7c-442c-971b-a95e437a3d23','Что мне хочется изменить в себе?',null),
                             ('e8033723-db99-4453-aad6-e3c2b3a997a3','14b225d6-6a7c-442c-971b-a95e437a3d23','Какой был мой последний сознательный выход из зоны комфорта? Понравился ли мне этот опыт?',null),
                             ('bab39075-9f4b-47f9-bd7c-621bd12869c8','14b225d6-6a7c-442c-971b-a95e437a3d23','Кого я недавно порадовал? Что это было?',null),
                             ('da1ef410-615e-4c15-8431-097a11782ee3','14b225d6-6a7c-442c-971b-a95e437a3d23','Есть ли в моем характере негативные черты? Как они могут влиять на меня и окружающих?',null),
                             ('158130b4-8bfd-4813-9370-52be45c8d51e','14b225d6-6a7c-442c-971b-a95e437a3d23','Что сейчас для меня самое важное? Почему оно так важно для меня?',null),
                             ('ed70d381-40e7-4140-a832-35c057211712','14b225d6-6a7c-442c-971b-a95e437a3d23','Есть ли дело, которое мне точно нужно сделать, но я его откладываю? Станет ли мне лучше, если я займусь им и завершу его?',null),
                             ('7c2b8be3-a814-4088-a0e1-7b82d18d2f47','14b225d6-6a7c-442c-971b-a95e437a3d23','Какие были бы мои планы на этот день, если бы он был последним в моей жизни?',null),
                             ('6b9025f0-b8d8-44d0-b384-7ab44fae1d6c','14b225d6-6a7c-442c-971b-a95e437a3d23','Какое мое самое большое беспокойство? Что люди обычно делают, чтобы закрыть для себя подобное?',null),
                             ('97131860-9fef-49f6-835e-c2906bcd4689','14b225d6-6a7c-442c-971b-a95e437a3d23','Ставлю ли я перед собой цели? Нужно ли мне делать это больше?',null),
                             ('259a54da-fd10-4f38-8c92-18a692d0b52b','14b225d6-6a7c-442c-971b-a95e437a3d23','Какие действия я делаю регулярно, чтобы заботиться о себе? Что к ним можно добавить?',null),
                             ('7efe7247-dffd-4d6b-9660-d097704f11c0','14b225d6-6a7c-442c-971b-a95e437a3d23','Боюсь ли я чего-то? Как я могу обезопаситься от этого?',null),
                             ('5b212706-063b-4ba0-b21b-56e53ec79882','14b225d6-6a7c-442c-971b-a95e437a3d23','Много ли у меня желания вставать по утрам? Если нет, какие у этого причины?',null),
                             ('51e9d3da-2c40-4a0c-9365-25c5435c612a','14b225d6-6a7c-442c-971b-a95e437a3d23','Когда мне последний раз пришлось врать? Почему? Можно ли было сказать правду?',null),
                             ('33bf80d1-c5ad-442d-862f-88dc5b59ee3e','14b225d6-6a7c-442c-971b-a95e437a3d23','Кем я хочу стать, чего добиться? Есть ли у меня план для достижения этого?',null),
                             ('af408481-f615-4307-9c75-e5bcc9823f22','14b225d6-6a7c-442c-971b-a95e437a3d23','Какие свои достижения я принимаю как должное?',null),
                             ('b88e575c-8f42-4a02-8a02-8899a11641cd','14b225d6-6a7c-442c-971b-a95e437a3d23','На что я трачу большую часть времени? А на что мне хотелось бы тратить?',null),
                             ('5759f141-2a38-47b7-a330-be01032bf61a','14b225d6-6a7c-442c-971b-a95e437a3d23','Какие сферы моей жизни устроены хорошо? А какие не очень?',null),
                             ('3aec5244-c01e-45f9-a154-5ae622aac1af','14b225d6-6a7c-442c-971b-a95e437a3d23','Образ какого человека я создаю?',null),
                             ('bc48528b-8dac-4fc8-a0c5-c119a6ef416f','14b225d6-6a7c-442c-971b-a95e437a3d23','В начале моего дела, какие у меня были мечты? Следую ли я им?',null),
                             ('ecf7ba13-1116-44a4-a116-24a792081523','14b225d6-6a7c-442c-971b-a95e437a3d23','Как у меня дела? А на самом деле?',null),
                             ('fe80c86e-82a3-41da-bb5f-873fdf76f71f','14b225d6-6a7c-442c-971b-a95e437a3d23','Какие утренние ритуалы я бы хотел добавить в свою жизнь?',null);