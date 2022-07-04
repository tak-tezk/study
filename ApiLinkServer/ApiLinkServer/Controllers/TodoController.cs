using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using ApiLinkServer.Models;

namespace ApiLinkServer.Controllers {
    [Route("api/[controller]")]
    [ApiController]
    public class TodoController : ControllerBase {
        // Todoアイテムの初期データ。本来はデータベースなどから取得する。
        private static List<TodoItem> items = new List<TodoItem>() {
            new TodoItem() { Id = 1, Name = @"犬の散歩", IsDone = false, },
            new TodoItem() { Id = 2, Name = @"買い物", IsDone = true, },
            new TodoItem() { Id = 3, Name = @"本棚の修理", IsDone = false },
        };

        [HttpGet]
        public ActionResult<List<TodoItem>> GetAll()
            => items;

        [HttpGet("{id}", Name = "Todo")]
        public ActionResult<TodoItem> GetById(int id) {
            var item = items.Find(i => i.Id == id);
            if (item == null)
                return NotFound();
            return item;
        }
    }
}
