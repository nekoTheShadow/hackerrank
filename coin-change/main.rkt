#lang racket

(define (main)
  (define nm (map string->number (string-split (read-line))))
  (define n (list-ref nm 0))
  (define m (list-ref nm 1))
  (define c (list->vector (map string->number (string-split (read-line)))))
  (define memo (make-hash))

  (define (dp n i)
    (let/cc return
      (when (= n 0)
        (return 1))
      (when (= i m)
        (return 0))
      (define key (list n i))
      (when (hash-has-key? memo key)
        (return (hash-ref memo key)))

      (define coin (vector-ref c i))
      (define value
        (for/sum ([count (in-range (+ (quotient n coin) 1))]) (dp (- n (* coin count)) (+ i 1))))
      (hash-set! memo key value)
      (return value)))

  (displayln (dp n 0)))

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
  (test-main "./00_input.txt" "4\n")
  (test-main "./01_input.txt" "5\n")
  ; (test-main "./input1.txt" (file->string "./output1.txt"))
  )
