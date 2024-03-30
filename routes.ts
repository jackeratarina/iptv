import {  Router } from "https://deno.land/x/oak/mod.ts";
import { tvHandler } from "./controllers/iptv.ts";

const router = new Router();

router.get("/iptv/:channel", tvHandler);

export default router;
