package dbSelect

const Actor string = `select id
                           , FullName
                        from Actor;`

const CountActor string = `select count(*) as Actors from Actor`
