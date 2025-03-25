
export function isJsonString(value: any) {
    try {
        const toObj = JSON.parse(value);
        if (toObj && typeof toObj === 'object') {
        return true;
        }
    } catch {}
    return false;
}