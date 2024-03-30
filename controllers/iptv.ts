import { Status } from "https://deno.land/x/oak@14.2.0/deps.ts";
import { Context } from "https://deno.land/x/oak@14.2.0/mod.ts";
import { getQuery } from "https://deno.land/x/oak@14.2.0/helpers.ts";
import getData from "../common/genkey.js";
async function tvHandler(ctx: Context): Promise<void> {
  const channel = getQuery(ctx, { mergeParams: true })["channel"];
//   switch (channel) {
//     case "vl1":
//       break;
//     case "vl2":
//       break;
//     case "vl3":
//       break;
//     case "vl4":
//       break;
//   }
const result = await (await getData(channel)).json();
const link_play = result?.play_info?.data?.link_play;
  ctx.response.redirect(link_play);
//   ctx.response.status = Status.;
}
export { tvHandler };
