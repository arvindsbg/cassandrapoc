CREATE KEYSPACE blackjack
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  };


CREATE TABLE facts (
    GameID text,
    TS timestamp,
    SpinExtID text,
    amount decimal,
    requestid text,
    tid text,
accountid text,
    PRIMARY KEY ((accountid, requestid), tid, spinextid, gameid)
);



insert into facts (accountid, requestid, tid, spinextid, gameid, amount, ts ) values ('testaccount2', '2123', '2456', '2789', '2111', 29.99,toTimeStamp(toDate(now())));
