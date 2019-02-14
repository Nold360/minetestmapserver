package postgres


const getBlocksByMtimeQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.mtime > ?
order by b.mtime asc
limit ?
`

const getLastBlockQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.mtime = 0
and b.posx >= ?
and b.posy >= ?
and b.posz >= ?
order by b.posx asc, b.posy asc, b.posz asc, b.mtime asc
limit ?
`

const countBlocksQuery = `
select count(*) from blocks b where b.mtime >= ? and b.mtime <= ?
`

const getBlockQuery = `
select pos,data,mtime from blocks b
where b.posx = ?
and b.posy = ?
and b.posz = ?
`
