
var grid_size = 20;
var qizi_radius = 7;
var WHITE = 2;
var BLACK = 1;
function drawQizi(ctx, i,j, color) {
  ctx.beginPath();
  ctx.arc(i*grid_size,j*grid_size,qizi_radius,0,2*Math.PI);
  if (color == WHITE) {
    ctx.fillStyle = '#ffffff';
  } else {
    ctx.fillStyle = '#000000';
  }
  ctx.fill();
  ctx.stroke();
}
