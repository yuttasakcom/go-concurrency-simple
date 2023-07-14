# Go Concurrency Simple

## Race Condition

```
สภาวะการเกิด race condition เป็นปัญหาที่เกิดขึ้นในการพัฒนาซอฟต์แวร์
หรือระบบที่มีการใช้งานพร้อมกันของกระบวนการหลายๆ กระบวนการ
โดยที่กระบวนการเหล่านั้นมีการเข้าถึงหรือปรับเปลี่ยนข้อมูลที่ใช้ร่วมกัน
ในหน่วยความจำหรือทรัพยากรร่วมกัน

```

## Go Mantra

```
Go's mantra for concurrency is: Do not communicate by sharing memory;
instead, share memory by communicating.
In this article we will review a simple concurrent program called headcheck ,
that, given a list of URLs,
performs an HTTP HEAD request on each one and reports its results.

```

## Reference

- [Introducing the Go Race Detector](https://go.dev/blog/race-detector)
- [Data Race Detector](https://go.dev/doc/articles/race_detector)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
