export function goToPage(page) {
    // eslint-disable-next-line no-restricted-globals
    location.replace(location.origin + page)
}

export function reload() {
    // eslint-disable-next-line no-restricted-globals
    location.reload();
}
