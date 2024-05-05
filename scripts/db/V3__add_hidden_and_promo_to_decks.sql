alter table decks add (
    hidden bool not null default false,
    promo varchar(255) unique
);

create table unlocked_decks (
    client_id varchar(255),
    deck_id varchar(255),
    constraint unique index (client_id, deck_id)
)