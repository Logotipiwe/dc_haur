create database if not exists `haur`;
use `haur`;
drop table if exists questions;
drop table if exists decks;
create table decks
(
    id  varchar(255) not null primary key,
    name varchar(255) not null,
    description text
);
create table questions
(
    id              varchar(255) not null primary key,
    level           varchar(60) not null,
    deck_id         varchar(255) references decks,
    text            varchar(255) not null
);

INSERT INTO decks values ('1', 'Для пары', '');

INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, что самое сложное в том деле, которым я зарабатываю себе на жизнь?');
INSERT INTO questions values (UUID(),'Знакомство','1','Что было первое, что ты заметил(а) во мне?');
INSERT INTO questions values (UUID(),'Знакомство','1','Если бы ты покупал(а) мне подарок, зная только как я выгляжу, что бы это было?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, растения у меня дома цветут или умирают? Объясни почему.');
INSERT INTO questions values (UUID(),'Знакомство','1','Я выгляжу добрым(ой)? Почему, или почему нет?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как тебе кажется, у меня более творческий или аналитический склад ума? Почему?');
INSERT INTO questions values (UUID(),'Знакомство','1','Я похож(а) на человека, который сделал бы татуировку чьего-то имени? Почему?');
INSERT INTO questions values (UUID(),'Знакомство','1','Закончите предложение: Просто взглянув на тебя, я бы подумал(а)...');
INSERT INTO questions values (UUID(),'Знакомство','1','Что моя обувь говорит тебе обо мне?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, меня когда-нибудь увольняли с работы? Если да, то за что?');
INSERT INTO questions values (UUID(),'Знакомство','1','Каково было твое первое впечатление обо мне?');
INSERT INTO questions values (UUID(),'Знакомство','1','Сделай предположение обо мне');
INSERT INTO questions values (UUID(),'Знакомство','1','Что тебя во мне интригует?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, кем я хотел(а) стать в детстве?');
INSERT INTO questions values (UUID(),'Знакомство','1','Смотря на тебя, я сразу думаю о...');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, я был(а) популярен(на) в школе? Объясни почему?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, сколько штрафов я получал(а) в своей жизни?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, легко ли я влюбляюсь? Почему или почему нет?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, какой комплимент я слышу чаще всего?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, я когда-либо проверял(а) переписки партнера?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, какой предмет в школе был моим любимым? А какой я не любил(а)?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, я предпочитаю чай, или кофе? С сахаром, или без?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, когда я был ребёнком, кем я хотел стать?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как тебе кажется, я жаворонок или сова?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, я был(а) популярен(на) в школе? Объясни почему');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, я чаще опаздываю, или прихожу вовремя? Объясни почему');
INSERT INTO questions values (UUID(),'Знакомство','1','Я похож(а) на кого-то из твоих знакомых?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, на что я готов потратить много денег?');
INSERT INTO questions values (UUID(),'Знакомство','1','Ты обычно коннектишься с такими как я?');
INSERT INTO questions values (UUID(),'Знакомство','1','Пугает ли тебя что-то во мне? Почему или почему нет?');
INSERT INTO questions values (UUID(),'Знакомство','1','Какая строчка из любимых песен приходит тебе в голову прямо сейчас?');
INSERT INTO questions values (UUID(),'Знакомство','1','Запиши свою главную цель на следующий месяц. Сравни с целью собеседника.');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, я был(а) когда-то влюблен(а)?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, есть ли типаж людей, который мне нравится. Какой он?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, каким(ой) я был(а) в школе?');
INSERT INTO questions values (UUID(),'Знакомство','1','Как ты думаешь, что значит мое имя?');
INSERT INTO questions values (UUID(),'Знакомство','1','Что во мне ты заметил(а) в первую очередь?');
INSERT INTO questions values (UUID(),'Знакомство','1','Ты хотел(а) бы стать известным(ой)? Каким образом?');
INSERT INTO questions values (UUID(),'Знакомство','1','За что в твоей жизни ты чувствуешь себя больше всего благодарным(ой)?');
INSERT INTO questions values (UUID(),'Знакомство','1','Если бы ты мог(ла) изменить что-нибудь в том, как тебя воспитывали, что бы это было?');
INSERT INTO questions values (UUID(),'Знакомство','1','Из чего состоит твой идеальный день?');
INSERT INTO questions values (UUID(),'Знакомство','1','Когда ты в последний раз пел(а) для себя? А для кого-то другого?');
INSERT INTO questions values (UUID(),'Знакомство','1','Внимательно посмотрите друг другу в глаза и опишите глаза собеседника');

INSERT INTO questions values (UUID(),'Погружение','1','Что было твоим счастливейшим воспоминанием за последний год?');
INSERT INTO questions values (UUID(),'Погружение','1','Ты изменил(а) свое мнение о чем-нибудь в последнее время?');
INSERT INTO questions values (UUID(),'Погружение','1','Какое твое самое раннее воспоминание о счастье?');
INSERT INTO questions values (UUID(),'Погружение','1','От какой привычки ты отчаивался(лась) дольше всего?');
INSERT INTO questions values (UUID(),'Погружение','1','Ты лжешь себе о чем-нибудь?');
INSERT INTO questions values (UUID(),'Погружение','1','На какие вопросы ты пытаешься ответить в своей жизни прямо сейчас?');
INSERT INTO questions values (UUID(),'Погружение','1','Когда в последний раз ты удивлялся себе?');
INSERT INTO questions values (UUID(),'Погружение','1','Какое название ты бы дал(а) этой главе в своей жизни?');
INSERT INTO questions values (UUID(),'Погружение','1','Чего бы тебе сейчас хотелось больше всего?');
INSERT INTO questions values (UUID(),'Погружение','1','Закончи предложение: Незнакомцы описали бы меня, как ..., но только знаю, что я ...');
INSERT INTO questions values (UUID(),'Погружение','1','Какую самую сильную нефизическую боль ты переживал(а)?');
INSERT INTO questions values (UUID(),'Погружение','1','Был ли случай, когда незнакомец изменил твою жизнь?');
INSERT INTO questions values (UUID(),'Погружение','1','Как звали твою первую любовь и почему ты влюбился(лась) в него/нее?');
INSERT INTO questions values (UUID(),'Погружение','1','Как бы ты описал(а) ощущение влюбленности одним словом?');
INSERT INTO questions values (UUID(),'Погружение','1','Чего ты больше боишься: провала или успеха и почему?');
INSERT INTO questions values (UUID(),'Погружение','1','От какой мечты ты отказался(лась)?');
INSERT INTO questions values (UUID(),'Погружение','1','Если бы вы могли узнать кого-то в своей жизни на более глубоком уровне, кто бы это был и почему?');
INSERT INTO questions values (UUID(),'Погружение','1','Как зовут твою маму, и что самое прекрасное в ней?');
INSERT INTO questions values (UUID(),'Погружение','1','Какая часть твоей жизни работает хорошо, а какая страдает?');
INSERT INTO questions values (UUID(),'Погружение','1','Как ты можешь стать лучшей версией себя?');
INSERT INTO questions values (UUID(),'Погружение','1','Ты по кому-нибудь сейчас скучаешь? Как ты думаешь они скучают по тебе?');
INSERT INTO questions values (UUID(),'Погружение','1','Как ты думаешь, в какой фаст-фуд ресторан я скорее всего пойду? Что я скорее всего закажу?');
INSERT INTO questions values (UUID(),'Погружение','1','Чем я могу отталкивать людей?');
INSERT INTO questions values (UUID(),'Погружение','1','Когда тебя спрашивают как дела, как часто ты отвечаешь правдиво?');
INSERT INTO questions values (UUID(),'Погружение','1','Как твои дела? А на самом деле?');
INSERT INTO questions values (UUID(),'Погружение','1','Что самое необычное, что случалось с тобой?');
INSERT INTO questions values (UUID(),'Погружение','1','Как зовут твоего отца? Расскажи мне одну вещь о нем');
INSERT INTO questions values (UUID(),'Погружение','1','Если бы ты мог(ла) быть где угодно и делать что угодно - где бы ты был(а) и что бы ты делал(а)?');
INSERT INTO questions values (UUID(),'Погружение','1','Опиши свой идеальный день');
INSERT INTO questions values (UUID(),'Погружение','1','Какой совет ты бы дал(а) себе из прошлого?');
INSERT INTO questions values (UUID(),'Погружение','1','На сколько лет ты себя ощущаешь?');
INSERT INTO questions values (UUID(),'Погружение','1','Что самое неловкое случалось с тобой на свидании?');
INSERT INTO questions values (UUID(),'Погружение','1','Что бы ты мог(ла) сделать лучше в предыдущих отношениях?');
INSERT INTO questions values (UUID(),'Погружение','1','Есть ли у тебя образ, от которого ты бы хотел избавиться?');
INSERT INTO questions values (UUID(),'Погружение','1','Чему люди, которые тебя воспитывали, научили тебя о любви?');
INSERT INTO questions values (UUID(),'Погружение','1','Опиши свое идеальное свидание');
INSERT INTO questions values (UUID(),'Погружение','1','Чему ты научился(лась) в романтических отношениях? Чему ты научился(лась), когда был(а) один(на)?');
INSERT INTO questions values (UUID(),'Погружение','1','Какой вопрос ты бы хотел(а), чтобы тебя чаще спрашивали?');
INSERT INTO questions values (UUID(),'Погружение','1','Чем ты страстно увлекаешься?');
INSERT INTO questions values (UUID(),'Погружение','1','Чему в данный момент ты уделяешь недостаточно времени?');


INSERT INTO questions values (UUID(),'Рефлексия','1','После нашего разговора. какую книгу ты бы посоветовал(а) мне почитать?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что во мне тебя удивило?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь какая у меня суперспособность?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, какие наши самые важные сходства?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, какую одну вещь я могу сделать, чтобы значительно улучшить свою жизнь?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что было бы идеальным подарком для меня?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как бы ты описал(а) меня незнакомцу(ке)?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что мне нужно услышать прямо сейчас?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Какой урок ты вынес(ла) из нашего разговора?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Чем я могу тебе помочь?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, чего я боюсь больше всего?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, о чем я могу давать советы?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что тебе труднее всего понять во мне?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Если бы мы были музыкальной группой, как бы мы назывались?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Признайся в чем-то');
INSERT INTO questions values (UUID(),'Рефлексия','1','Какие части себя ты видишь во мне?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что помогает тебе раскрыться перед другим человеком?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что ты посоветуешь мне отпустить?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Чему этот разговор научил тебя о себе?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, какая моя самая определяющая черта характера?');
INSERT INTO questions values (UUID(),'Рефлексия','1','На какой вопрос тебе было сложнее всего ответить?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, почему мы встретились?');
INSERT INTO questions values (UUID(),'Рефлексия','1','После нашего разговора что ты будешь помнить обо мне?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь в чем моя слабость?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как наши личности дополняют друг друга?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, что я должен знать о себе, чего, возможно, я не знаю?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как бы ты почувствовал(а) себя ближе ко мне?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Опиши как ты чувствуешь себя прямо сейчас одним словом');
INSERT INTO questions values (UUID(),'Рефлексия','1','Веришь ли ты, что у каждого человека есть призвание? Если это так, как ты думаешь, я нашел(ла) свое?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что мы можем создать вместе?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Какой мой ответ дал тебе инсайт о себе?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что не в моей внешности привлекает тебя больше всего?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Совпали ли первые впечатления от меня с тем, как я раскрывался(лась) в течение нашего разговора?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что тебя больше всего восхищает во мне?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как бы ты описал меня незнакомцу?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Чему я могу научить тебя?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что бы ты хотел(а) узнать от меня?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Что во мне больше всего удивило тебя?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Чем я могу тебе помочь?');
INSERT INTO questions values (UUID(),'Рефлексия','1','О ком ты часто  думаешь в последнее время? Придумай, как дать им знать, что ты думаешь о них');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как бы ты описал(а) наш разговор в двух словах?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Чему в твоей жизни сегодня не поверил бы ты из прошлого?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как наши характеры дополняют друг друга?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как ты думаешь, совпадает ли то, как видят тебя другие люди с тем, какой(а)я ты есть на самом деле?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Как я могу добавить 1% счастья в твою жизнь?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Скажи три правдивых высказывания о нас. Например, "Сейчас мы оба чувствуем..."');
INSERT INTO questions values (UUID(),'Рефлексия','1','Когда ты в последний раз плакал(а) перед другим человеком? А в одиночестве?');
INSERT INTO questions values (UUID(),'Рефлексия','1','Каким способом лучше всего показать свою любовь к кому-то?');
