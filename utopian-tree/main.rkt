#lang racket

(define (main)
  (define t (string->number (read-line)))
  (for ([_i (in-range t)])
    (define n (string->number (read-line)))
    (displayln (utopia-tree n))))

(define (utopia-tree n)
  (let loop ([count 0]
             [tree 1]
             [season #t])
    (if (= count n)
        tree
        (if season
            (loop (+ count 1) (* tree 2) (not season))
            (loop (+ count 1) (+ tree 1) (not season))))))

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

  (test-main "./input.txt" (file->string "./output.txt")))
