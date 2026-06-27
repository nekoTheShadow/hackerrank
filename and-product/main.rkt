#lang racket

(define (and-product a b)
  (let loop ([x a]
             [y b]
             [count 0])
    (if (= x y)
        (arithmetic-shift x count)
        (loop (arithmetic-shift x -1) (arithmetic-shift y -1) (+ count 1)))))

(define (main)
  (define n (string->number (read-line)))
  (for ([_i (in-range n)])
    (define ab (map string->number (string-split (read-line))))
    (define a (first ab))
    (define b (second ab))
    (displayln (and-product a b))))

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
  (test-main "./input0.txt" (file->string "./output0.txt"))
  (test-main "./input1.txt" (file->string "./output1.txt"))

  ; テストを実装する
  ; 例：
  ; (test-main "./input0.txt" "1\n")
  ; (test-main "./input1.txt" (file->string "./output1.txt"))
  )
