let res = ""

res += "INSERT INTO `decks` (`id`, `language_code`, `name`, `emoji`, `description`, `labels`, `vector_image_id`, `hidden`, `promo`) VALUES ";

res += Array.from(document.querySelector('[data-block-id=de41a85a-3671-43a2-800a-1ce43ef93c59]')
    .querySelectorAll('.notion-table-view-row'))
    .map(row => {
        const cellsObjs = Array.from(row.querySelectorAll('.notion-table-view-cell'))
        const cells = cellsObjs.map(cell => cell.innerText)
        const emoji = cellsObjs[3].querySelector('.notion-emoji')
        return `('${cells[0]}', '${cells[1].substring(0,2).toUpperCase()}', '${cells[2]}', ${emoji ? `'${emoji.alt}'` : 'null'}, '${cells[4]}', 'good to start', '${cells[6]}', 0, null)`
    }).join(", ")

res += ";"
res += "INSERT INTO `levels` (`id`, `deck_id`, `level_order`, `name`, `emoji`, `color_start`, `color_end`, `color_button`) VALUES "

res += Array.from(document.querySelectorAll('.notion-collection_view-block')[4]
    .querySelectorAll('.notion-table-view-row'))
    .map(row => {
        const cellsObjs = Array.from(row.querySelectorAll('.notion-table-view-cell'))
        const cells = cellsObjs.map(cell => cell.innerText)
        const emoji = cellsObjs[5].querySelector('.notion-emoji')
        return `('${cells[1]}', '${cells[2]}', ${cells[3]}, '${cells[4]}', ${emoji ? `'${emoji.alt}'` : 'null'}, '${cells[6]}', '${cells[7]}', '${cells[8]}')`
    }).join(", ");

res += ";"
res += "INSERT INTO `questions` (`id`, `level_id`, `text`, `additional_text`) VALUES ";

res += Array.from(document.querySelectorAll('.notion-collection_view-block')[8]
    .querySelectorAll('.notion-table-view-row'))
    .map(row => {
        const cellsObjs = Array.from(row.querySelectorAll('.notion-table-view-cell'))
        const cells = cellsObjs.map(cell => cell.innerText)
        const additional = cells[3]
        return `('${cells[0]}', '${cells[1]}', '${cells[2]}', ${additional ? `'${additional}'` : 'null'})`;
    }).join(", ")