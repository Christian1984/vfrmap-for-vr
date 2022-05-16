import md5 from "md5";

const SENDER_ID = md5(Math.floor(Math.random() * Number.MAX_VALUE).toString());

export default SENDER_ID;