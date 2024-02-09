function toggleMonster(nameID) {
    let popups = document.getElementsByClassName("popupMonster");
    for (let i = 0; i < popups.length; i++) {
        let popup = popups[i];
        if (popup.dataset.monsterId === nameID) {
            popup.classList.toggle("open");
            console.log(nameID);
            return;
        }
    }
    console.error("Popup not found for artist with id: " + nameID);
}
