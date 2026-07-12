#lang racket

(define (main)
  (define NP (map string->number (string-split (read-line))))
  (define N (list-ref NP 0))
  (define P (list-ref NP 1))

  (define parents
    (for/vector ([i (in-range N)])
      i))
  (for ([_i (in-range P)])
    (define XY (map string->number (string-split (read-line))))
    (define X (list-ref XY 0))
    (define Y (list-ref XY 1))
    (union parents X Y))

  (define ht (make-hash))
  (for ([i (in-range N)])
    (hash-update! ht (find parents i) add1 0))

  (define ans (quotient (* N (- N 1)) 2))
  (for ([v (hash-values ht)])
    (define w (quotient (* v (- v 1)) 2))
    (set! ans (- ans w)))
  (displayln ans))

(define (find parents x)
  (let/ec return
    (when (= x (vector-ref parents x))
      (return x))

    (vector-set! parents x (find parents (vector-ref parents x)))
    (return (vector-ref parents x))))

(define (union parents x y)
  (define a (find parents x))
  (define b (find parents y))
  (unless (= a b)
    (vector-set! parents a b)))

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
  (test-main "./00_in.txt" "6\n")
  (test-main "./01_in.txt" "5\n")
  ; (test-main "./input1.txt" (file->string "./output1.txt"))
  )
