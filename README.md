Bowling Kata - event souring
===

## @8f3fd6a

### 成果：

Bowling struct 的 command 與 event 模組分離完成。

Bowling.changes 能夠不斷的紀錄 event，接下來可理用紀錄的 event 重播 Bowling 的完整歷程。

### 下一步：

部分邏輯與事件處理行為混在一起，如果沒有按照順序執行將導致錯誤。

由於目前使用 cache 紀錄物件，且物件不會被同時操作，所以沒有錯誤發生。

Aggregate 應該負責處理邏輯而非儲存 Aggregate State。

`Command -> Function -> Event`

`Event + State -> Function -> new State`
