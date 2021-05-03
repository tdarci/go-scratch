function buttonClick(displayField) {
    console.log("button clicked")

    fetch('/time').
        then(response => response.json()).
        then(data => {
            console.log(data)
            const e = document.getElementById(displayField)
            e.innerHTML = data.current_time
        }
    )

}