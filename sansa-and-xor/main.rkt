#lang racket

(define (main)
  (define t (string->number (read-line)))
  (for ([_i (in-range t)])
    (define n (string->number (read-line)))
    (define arr (map string->number (string-split (read-line))))
    (displayln (solve n arr))))

(define (solve n arr)
  (let/ec return
    (when (even? n)
      (return 0))

    (for/fold ([acc 0]) ([(v i) (in-indexed arr)])
      (if (even? i)
          (bitwise-xor acc v)
          acc))))

(module+ main
  (main))

(module+ test
  (require rackunit)

  (define (test-main filename expected)
    (define (run)
      (call-with-output-string run-with-out))

    (define (run-with-out stdout)
      (call-with-input-file filename (lambda (stdin) (run-with-in-out stdin stdout))))

    (define (run-with-in-out stdin stdout)
      (parameterize ([current-input-port stdin]
                     [current-output-port stdout])
        (main)))

    (check-equal? (run) expected))

  ; テストを実装する
  ; 例：
  (test-main "./input1.txt" "110\n48\n")
  (test-main "./input0.txt" "2\n0\n"))
