namespace date {
  export function dateToYMD(date: Date) {
    const d = date.getDate();
    const m = date.getMonth() + 1;
    const y = date.getFullYear();
    return `${y}-${m <= 9 ? `0${m}` : m}-${d <= 9 ? `0${d}` : d}`;
  }

  export function dateFromYMD(s: string) {
    const d = new Date(s);
    return new Date(d.getTime() + d.getTimezoneOffset() * 60000);
  }

  export function dow(i: number) {
    switch (i) {
      case 0:
        return "Sun";
      case 1:
        return "Mon";
      case 2:
        return "Tue";
      case 3:
        return "Wed";
      case 4:
        return "Thu";
      case 5:
        return "Fri";
      case 6:
        return "Sat";
      default:
        return "???";
    }
  }

  export function toDateString(d: Date) {
    return d.toLocaleDateString();
  }

  export function toTimeString(d: Date) {
    return d.toLocaleTimeString().slice(0, 8);
  }

  export function toDateTimeString(d: Date) {
    return `${toDateString(d)} ${toTimeString(d)}`;
  }

  const tzOffset = new Date().getTimezoneOffset() * 60000;

  export function utcDate(s: string) {
    return new Date(Date.parse(s) + tzOffset);
  }
}
